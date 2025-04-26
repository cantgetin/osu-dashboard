package config

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"path/filepath"
	"runtime"
	"time"
)

type Config struct {
	AppName  string `env:"APP_NAME" envDefault:"playcount-monitor-backend"`
	HTTPAddr string `env:"HTTP_ADDR" envDefault:":8080"`

	PgDSN          string        `env:"PG_DSN" envDefault:"postgresql://pmb-db:5432/db?user=db&password=db"`
	PgMaxOpenConn  int           `env:"PG_MAX_OPEN_CONN" envDefault:"5"`
	PgIdleConn     int           `env:"PG_MAX_IDLE_CONN" envDefault:"5"`
	PgPingInterval time.Duration `env:"PG_PING_INTERVAL" envDefault:"5s"`
	PgPassword     string        `env:"POSTGRES_PASSWORD" envDefault:""`

	TrackingTimeout  time.Duration `env:"TRACKING_TIMEOUT" envDefault:"30m"`
	TrackingInterval time.Duration `env:"TRACKING_INTERVAL" envDefault:"24h"`

	CleaningTimeout  time.Duration `env:"CLEANING_TIMEOUT" envDefault:"30m"`
	CleaningInterval time.Duration `env:"CLEANING_INTERVAL" envDefault:"24h"`

	OsuAPIClientID     string `env:"OSU_API_CLIENT_ID" envDefault:""`
	OsuAPIClientSecret string `env:"OSU_API_CLIENT_SECRET" envDefault:""`
	OsuAPIRedirectURI  string `env:"OSU_API_REDIRECT_URI" envDefault:""`

	OsuAPIHost   string `env:"OSU_API_HOST" envDefault:"https://osu.ppy.sh/api/v2"`
	OsuOAuthHost string `env:"OSU_OAUTH_HOST" envDefault:"https://osu.ppy.sh/oauth/token"`

	RunIntegrationTest bool `env:"RUN_INTEGRATION_TEST" envDefault:"false"`

	IntegrationTestPgDSN  string `env:"INTEGRATION_TEST_PG_DSN" envDefault:"postgresql://db:5467/db?user=db&password=db"`
	IntegrationTestPgPort string `env:"INTEGRATION_TEST_PG_PORT" envDefault:"5467"`

	IntegrationTestHTTPAddr string `env:"INTEGRATION_TEST_HTTP_ADDR" envDefault:"127.0.0.1:8155"`
	IntegrationTestHTTPPort string `env:"INTEGRATION_TEST_HTTP_PORT" envDefault:"8155"`
}

func LoadConfig(envFileName string) (*Config, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("failed to get current file information")
	}

	baseDir := filepath.Dir(filename)
	envDir := filepath.Join(baseDir, "..", "..", envFileName) // .env file path

	if err := godotenv.Load(envDir); err != nil {
		return nil, err
	}

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
