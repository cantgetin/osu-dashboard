package main

import (
	"context"
	"github.com/caarlos0/env"
	log "github.com/sirupsen/logrus"
	"playcount-monitor-backend/internal/app"
	"playcount-monitor-backend/internal/config"
)

func main() {
	cfg := &config.Config{}

	if err := env.Parse(cfg); err != nil {
		log.Fatalf("failed to retrieve env variables, %v", err)
	}

	lg := log.New()

	ctx, _ := context.WithCancel(context.Background())

	if err := app.RunTrackingWorker(ctx, cfg, lg); err != nil {
		log.Fatalf("failed to start app, %v", err)
	}
}
