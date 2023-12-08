package config

type Config struct {
	AppName  string `env:"APP_NAME" envDefault:"playcount-monitor-backend"`
	PgDSN    string `env:"PG_DSN" envDefault:"postgresql://localhost:5468/db?user=db&password=db"`
	HTTPAddr string `env:"HTTP_ADDR" envDefault:":8080"`
}
