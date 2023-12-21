package main

import (
	"playcount-monitor-backend/internal/app"
	"playcount-monitor-backend/internal/config"

	"github.com/caarlos0/env"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfg := &config.Config{}

	if err := env.Parse(cfg); err != nil {
		log.Fatalf("failed to retrieve env variables, %v", err)
	}

	lg := log.New()

	if err := app.Run(cfg, lg); err != nil {
		log.Fatalf("failed to start app, %v", err)
	}
}
