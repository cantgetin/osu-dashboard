package bootstrap

import (
	"fmt"
	"osu-dashboard/internal/config"
	"osu-dashboard/internal/database/txmanager"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/ds248a/closer"
	migrate "github.com/rubenv/sql-migrate"
	log "github.com/sirupsen/logrus"

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

	// Switch PgDSN password to secret password from .env
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

func ApplyMigrations(gdb *gorm.DB) error {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("failed to get current file information")
	}

	baseDir := filepath.Dir(filename)
	migrationsDir := filepath.Join(baseDir, "..", "..", "migrations")

	migrations := migrate.FileMigrationSource{
		Dir: migrationsDir,
	}

	db, err := gdb.DB()
	if err != nil {
		return err
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		return err
	}
	fmt.Printf("Applied %d migrations!\n", n)

	return nil
}

func ConnectTxManager(db *gorm.DB, lg *log.Logger) txmanager.TxManager {
	txm := txmanager.New(db, lg)
	return txm
}
