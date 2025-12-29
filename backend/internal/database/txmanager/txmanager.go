package txmanager

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"runtime/debug"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"gorm.io/gorm"
)

var (
	ErrTxDeadlock = errors.New("transaction deadlock occurred")
	ErrTxPanicked = errors.New("transaction failed")
	ErrTxNoRetry  = errors.New("tx cannot be retried")
)

type Tx interface {
	DB() *gorm.DB
}

type GormTx struct {
	db *gorm.DB
}

func (g *GormTx) DB() *gorm.DB {
	return g.db
}

type Effector func(ctx context.Context, tx Tx) error

type txConfig struct {
	level       sql.IsolationLevel
	maxAttempts int
	readOnly    bool
}

type TxConfigurator func(cfg *txConfig)

type TxManager interface {
	ReadWrite(ctx context.Context, effector Effector, configurators ...TxConfigurator) error
	ReadOnly(ctx context.Context, effector Effector, configurators ...TxConfigurator) error
}

type GormTxManager struct {
	db      *gorm.DB
	lg      *log.Logger
	metrics *prometheus.HistogramVec
}

func New(gdb *gorm.DB, metrics *prometheus.HistogramVec, lg *log.Logger) *GormTxManager {
	return &GormTxManager{
		db:      gdb,
		metrics: metrics,
		lg:      lg,
	}
}

func (tm *GormTxManager) ReadWrite(ctx context.Context, effector Effector, configurators ...TxConfigurator) error {
	cfg := &txConfig{
		level:       sql.LevelSerializable,
		maxAttempts: 1,
		readOnly:    false,
	}

	for _, configurator := range configurators {
		configurator(cfg)
	}

	var err error
	for attempt := 0; attempt < cfg.maxAttempts; attempt++ {
		err = nil

		if err = tm.execUnderLock(ctx, effector, cfg); err != nil {
			if errors.Is(err, ErrTxPanicked) || errors.Is(err, ErrTxNoRetry) {
				tm.lg.Errorf("read-write tx attempt %d failed: %v - no retry", attempt, err)
				break
			}

			tm.lg.Errorf("read-write tx attempt %d failed: %v", attempt, err)
		} else {
			break
		}
	}

	return err
}

func (tm *GormTxManager) ReadOnly(ctx context.Context, effector Effector, configurators ...TxConfigurator) error {
	cfg := &txConfig{
		level:       sql.LevelSerializable,
		readOnly:    true,
		maxAttempts: 1,
	}

	for _, configurator := range configurators {
		configurator(cfg)
	}

	var err error
	for attempt := 0; attempt < cfg.maxAttempts; attempt++ {
		err = nil

		if err = tm.execUnderLock(ctx, effector, cfg); err != nil {
			if errors.Is(err, ErrTxPanicked) || errors.Is(err, ErrTxNoRetry) {
				tm.lg.Errorf("read only tx attempt %d failed: %v - no retry", attempt, err)
				break
			}

			tm.lg.Errorf("read only tx attempt %d failed: %v", attempt, err)
		} else {
			break
		}
	}

	return err
}

func (tm *GormTxManager) execUnderLock(
	ctx context.Context,
	effector Effector,
	cfg *txConfig,
) (err error) {
	start := time.Now()
	defer func() {
		var t string
		if cfg.readOnly {
			t = "read-only"
		} else {
			t = "read-write"
		}

		tm.metrics.
			WithLabelValues(cfg.level.String(), t).
			Observe(float64(time.Since(start).Milliseconds()))
	}()

	tx := tm.db.Begin(&sql.TxOptions{
		Isolation: cfg.level,
		ReadOnly:  cfg.readOnly,
	})

	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tm.lg.Errorf("panic occurred inside transaction: %v. %s Rolling back", r, string(debug.Stack()))

			if rollbackErr := tx.Rollback().Error; rollbackErr != nil {
				tm.lg.Errorf("could not rollback: %s", rollbackErr.Error())
			}

			err = ErrTxPanicked
		}
	}()

	gormTx := &GormTx{db: tx}
	if err = effector(ctx, gormTx); err != nil {
		tm.lg.
			WithField("read-only", cfg.readOnly).
			WithField("isolation", cfg.level.String()).
			Errorf("tx failed: %v", err)

		if strings.Contains(strings.ToLower(err.Error()), "deadlock") {
			err = fmt.Errorf(
				"%w: read-only: %v, isolation: %s, base error: %s",
				ErrTxDeadlock,
				cfg.readOnly,
				cfg.level.String(),
				err.Error(),
			)
		}

		if rollbackErr := tx.Rollback().Error; rollbackErr != nil {
			tm.lg.WithField("read-only", cfg.readOnly).
				WithField("isolation", cfg.level.String()).
				Errorf("could not rollback: %s", rollbackErr.Error())
		}

		return err
	}

	return tx.Commit().Error
}
