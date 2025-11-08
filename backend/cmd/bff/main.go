package main

import (
	"context"
	"osu-dashboard/internal/app"
	"osu-dashboard/internal/config"

	log "github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatalf("failed to load config, %v", err)
	}

	lg := log.New()

	ctx := context.Background()

	if err := app.Run(ctx, cfg, lg); err != nil {
		log.Fatalf("failed to start app, %v", err)
	}
}
