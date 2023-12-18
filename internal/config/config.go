package config

import "time"

type Config struct {
	AppName  string `env:"APP_NAME" envDefault:"playcount-monitor-backend"`
	HTTPAddr string `env:"HTTP_ADDR" envDefault:":8080"`

	PgDSN          string        `env:"PG_DSN" envDefault:"postgresql://localhost:5468/db?user=db&password=db"`
	PgMaxOpenConn  int           `env:"PG_MAX_OPEN_CONN" envDefault:"5"`
	PgIdleConn     int           `env:"PG_MAX_IDLE_CONN" envDefault:"5"`
	PgPingInterval time.Duration `env:"PG_PING_INTERVAL" envDefault:"5s"`
}
