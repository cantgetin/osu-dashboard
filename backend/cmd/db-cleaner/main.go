package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"playcount-monitor-backend/internal/app"
	"playcount-monitor-backend/internal/config"
)

func main() {
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatalf("failed to load config, %v", err)
	}

	lg := log.New()

	ctx := context.Background()

	if err := app.RunDBCleaner(ctx, cfg, lg); err != nil {
		log.Fatalf("failed to start cleaning worker app, %v", err)
	}
}
