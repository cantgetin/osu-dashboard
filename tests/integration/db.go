package integration

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	migrate "github.com/rubenv/sql-migrate"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"playcount-monitor-backend/internal/config"
	"testing"
)

type CloseFn func() error

func InitDB(t *testing.T, pool *dockertest.Pool, cfg *config.Config) CloseFn {
	var db *sql.DB

	cfg.PgDSN = "postgresql://localhost:5432/db?user=db&password=db&sslmode=disable"
	if retryErr := pool.Retry(func() error {
		var err error
		db, err = sql.Open("postgres", cfg.PgDSN)
		if err != nil {
			return err
		}
		return db.Ping()
	}); retryErr != nil {
		log.Fatalf("Could not connect to database: %s", retryErr)
	}

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	migrations := &migrate.FileMigrationSource{
		Dir: filepath.Join(dir, "..", "migrations"),
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		t.Logf("executed %d migrations", n)
		t.Fatalf("Could not run the 'UP' migrations: %v", err)
	} else {
		if n < 1 {
			t.Fatal("should be at least 1 migration")
		}
		t.Logf("executed %d migrations", n)
	}

	closer := func() error {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
		return nil
	}

	return closer
}
