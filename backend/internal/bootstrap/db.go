package bootstrap

import (
	"fmt"
	"github.com/ds248a/closer"
	log "github.com/sirupsen/logrus"
	"playcount-monitor-backend/internal/config"
	"playcount-monitor-backend/internal/database/txmanager"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	logger := log.Logger{}
	gormConfig := &gorm.Config{
		DisableAutomaticPing: true,
		Logger: gormLogger.New(
			&logger,
			gormLogger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  gormLogger.Silent,
				IgnoreRecordNotFoundError: true,
				Colorful:                  false,
			},
		)}

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
