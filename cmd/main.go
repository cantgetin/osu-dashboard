package main

import (
	"github.com/caarlos0/env"
	"log"
	"playcount-monitor-backend/internal/app"
	"playcount-monitor-backend/internal/config"
)

func main() {
	cfg := &config.Config{}

	if err := env.Parse(cfg); err != nil {
		log.Fatalf("failed to retrieve env variables, %v", err)
	}

	lg := log.New(log.Default().Writer(), cfg.AppName, log.LstdFlags)

	if err := app.Run(cfg, lg); err != nil {
		log.Fatalf("error running grpc server, %v", err)
	}
}
