package integration

import (
	"os"
	"osu-dashboard/internal/bootstrap"
	"osu-dashboard/internal/config"
	"path/filepath"
	"testing"

	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	migrate "github.com/rubenv/sql-migrate"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CloseFn func() error

func InitDB(t *testing.T, pool *dockertest.Pool, cfg *config.Config) (*gorm.DB, CloseFn) {
	var gdb *gorm.DB

	if retryErr := pool.Retry(func() error {
		var err error
		gdb, err = bootstrap.InitDB(cfg)
		if err != nil {
			return err
		}
		if db, err := gdb.DB(); err != nil {
			return err
		} else {
			return db.Ping()
		}
	}); retryErr != nil {
		log.Fatalf("Could not connect to database: %s", retryErr)
	}

	db, err := gdb.DB()
	if err != nil {
		t.Fatal(err)
	}

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	migrations := &migrate.FileMigrationSource{
		Dir: filepath.Join(dir, "..", "migrations"),
	}

	migrationsCount, err := countMigrations("./migrations")
	if err != nil {
		t.Fatalf("Error counting migrations: %v", err)
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		t.Logf("executed %d migrations", n)
		t.Fatalf("Could not run the 'UP' migrations: %v", err)
	} else {
		if n < migrationsCount {
			t.Fatalf("should be at least %v migrations", migrationsCount)
		}
		t.Logf("executed %d migrations", n)
	}

	closer := func() error {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
		return nil
	}

	return gdb, closer
}

func countMigrations(folderPath string) (int, error) {
	files, err := filepath.Glob(filepath.Join(folderPath, "*"))
	if err != nil {
		return 0, err
	}
	return len(files), nil
}
