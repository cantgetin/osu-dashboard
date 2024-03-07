package main

import (
	"context"
	"playcount-monitor-backend/internal/app"
	"playcount-monitor-backend/internal/config"

	log "github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatalf("failed to load config, %v", err)
	}

	lg := log.New()

	ctx, _ := context.WithCancel(context.Background())

	if err := app.Run(ctx, cfg, lg); err != nil {
		log.Fatalf("failed to start app, %v", err)
	}
}
