package integration

import (
	"fmt"
	"github.com/ory/dockertest/v3"
	dc "github.com/ory/dockertest/v3/docker"
	"osu-dashboard/internal/config"
	"testing"
)

func Start(t *testing.T, cfg *config.Config) (*dockertest.Pool, CloseFn) {
	t.Helper()

	pgIntegrationTestExternalPort := cfg.IntegrationTestPgPort
	pgInternalPort := "5432/tcp"

	pool, err := dockertest.NewPool("")
	if err != nil {
		t.Fatal(err)
	}

	postgres, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "15.2-alpine",
		Env: []string{
			"POSTGRES_USER=db",
			"POSTGRES_PASSWORD=" + cfg.PgPassword,
			"POSTGRES_DB=db",
		},
		ExposedPorts: []string{pgInternalPort},
		PortBindings: map[dc.Port][]dc.PortBinding{
			dc.Port(pgInternalPort): {{HostIP: "", HostPort: pgIntegrationTestExternalPort}},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	cfg.PgDSN = fmt.Sprintf("postgresql://localhost:%s/db?user=db&password=db&sslmode=disable",
		pgIntegrationTestExternalPort)

	closer := func() error {
		if err := pool.Purge(postgres); err != nil {
			t.Error(err)
		}
		return nil
	}

	return pool, closer
}
