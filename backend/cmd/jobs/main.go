package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"osu-dashboard/internal/app"
	"osu-dashboard/internal/config"
)

func main() {
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatalf("failed to load config, %v", err)
	}

	lg := log.New()

	ctx := context.Background()

	if err := app.RunJobs(ctx, cfg, lg); err != nil {
		log.Fatalf("failed to start jobs app, %v", err)
	}
}
