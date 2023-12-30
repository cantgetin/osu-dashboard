package integration

import (
	"fmt"
	"github.com/ory/dockertest/v3"
	dc "github.com/ory/dockertest/v3/docker"
	"playcount-monitor-backend/internal/config"
	"testing"
)

func Start(t *testing.T, cfg *config.Config) (*dockertest.Pool, CloseFn) {
	t.Helper()

	pool, err := dockertest.NewPool("")
	if err != nil {
		t.Fatal(err)
	}

	postgres, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "15.2-alpine",
		Env: []string{
			"POSTGRES_USER=db",
			"POSTGRES_PASSWORD=db",
			"POSTGRES_DB=db",
		},
		ExposedPorts: []string{"5432/tcp"}, // Specify the container's exposed port
		PortBindings: map[dc.Port][]dc.PortBinding{
			"5432/tcp": {{HostIP: "", HostPort: "5432"}}, // Beatmapset the exposed port to the desired host port
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	postgresPort := postgres.GetPort("5432/tcp")
	cfg.PgDSN = fmt.Sprintf("postgres://db:db@db:%s/?sslmode=disable", postgresPort)

	closer := func() error {
		if err := pool.Purge(postgres); err != nil {
			t.Error(err)
		}
		return nil
	}

	return pool, closer
}
