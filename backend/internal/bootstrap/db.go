package bootstrap

import (
	"fmt"
	"github.com/ds248a/closer"
	migrate "github.com/rubenv/sql-migrate"
	log "github.com/sirupsen/logrus"
	"path/filepath"
	"osu-dashboard/internal/config"
	"osu-dashboard/internal/database/txmanager"
	"runtime"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	gormConfig := &gorm.Config{
		DisableAutomaticPing: true,
		Logger: gormLogger.New(
			&log.Logger{},
			gormLogger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  gormLogger.Info,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		)}

	// switch pgDSN to secret password from .env
	cfg.PgDSN = strings.Replace(cfg.PgDSN, "password=db", "password="+cfg.PgPassword, 1)

	db, err := gorm.Open(postgres.Open(cfg.PgDSN), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("can't connect to pg instance, %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(cfg.PgIdleConn)
	sqlDB.SetMaxOpenConns(cfg.PgMaxOpenConn)

	go func() {
		t := time.NewTicker(cfg.PgPingInterval)

		for range t.C {
			if err := sqlDB.Ping(); err != nil {
				log.Errorf("error ping %v", err)
			}
		}
	}()

	closer.Add(func() {
		_ = sqlDB.Close
	})

	return db, nil
}

// TODO: look at this later, applying existing migrations causing it to fail
func ApplyMigrations(gdb *gorm.DB) error {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("failed to get current file information")
	}

	// Get the directory containing the current file
	baseDir := filepath.Dir(filename)
	migrationsDir := filepath.Join(baseDir, "..", "..", "migrations")

	// Create a migration source with the absolute path
	migrations := migrate.FileMigrationSource{
		Dir: migrationsDir,
	}

	db, err := gdb.DB()
	if err != nil {
		return err
	}

	// Apply migrations
	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		return err
	}
	fmt.Printf("Applied %d migrations!\n", n)

	return nil
}

func ConnectTxManager(ns string, wait time.Duration, db *gorm.DB, lg *log.Logger) txmanager.TxManager {
	m := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: ns,
		Subsystem: "database",
		Name:      "query_time_milliseconds",
		Help:      "query duration",
		Buckets:   []float64{1, 2.5, 5, 10, 25, 50, 100, 500, 1000},
	}, []string{"query", "query2"})

	txm := txmanager.New(db, m, lg)
	return txm
}
