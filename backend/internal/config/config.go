package config

import "time"

type Config struct {
	AppName        string        `env:"APP_NAME" envDefault:"playcount-monitor-backend"`
	HTTPAddr       string        `env:"HTTP_ADDR" envDefault:":8080"`
	PgDSN          string        `env:"PG_DSN" envDefault:"postgresql://db:5432/db?user=db&password=db"`
	PgMaxOpenConn  int           `env:"PG_MAX_OPEN_CONN" envDefault:"5"`
	PgIdleConn     int           `env:"PG_MAX_IDLE_CONN" envDefault:"5"`
	PgPingInterval time.Duration `env:"PG_PING_INTERVAL" envDefault:"5s"`

	TrackingTimeout  time.Duration `env:"TRACKING_TIMEOUT" envDefault:"5m"`
	TrackingInterval time.Duration `env:"TRACKING_INTERVAL" envDefault:"24h"`

	OsuAPIClientID     string `env:"OSU_API_CLIENT_ID" envDefault:""`
	OsuAPIClientSecret string `env:"OSU_API_CLIENT_SECRET" envDefault:""`

	OsuAPIHost   string `env:"OSU_API_HOST" envDefault:"https://osu.ppy.sh/api/v2"`
	OsuOAuthHost string `env:"OSU_OAUTH_HOST" envDefault:"https://osu.ppy.sh/oauth/token"`
}
