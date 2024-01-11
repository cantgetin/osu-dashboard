package app

import (
	"context"
	"fmt"
	"github.com/ds248a/closer"
	log "github.com/sirupsen/logrus"
	"playcount-monitor-backend/internal/app/trackingworkerapi"
	"playcount-monitor-backend/internal/bootstrap"
	"playcount-monitor-backend/internal/config"
	"playcount-monitor-backend/internal/database/repository/beatmaprepository"
	"playcount-monitor-backend/internal/database/repository/mapsetrepository"
	"playcount-monitor-backend/internal/database/repository/trackingrepository"
	"playcount-monitor-backend/internal/database/repository/userrepository"
	"playcount-monitor-backend/internal/usecase/track"
	"time"
)

func RunTrackingWorker(
	ctx context.Context,
	cfg *config.Config,
	lg *log.Logger,
) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	closer.Add(cancel)

	db, err := bootstrap.InitDB(cfg)
	if err != nil {
		return fmt.Errorf("failed to init db: %w", err)
	}

	const waitForConnection = 5 * time.Second
	txm := bootstrap.ConnectTxManager("tracking-worker", waitForConnection, db, lg)

	// init repos
	userRepo, err := userrepository.New(cfg, lg)
	mapsetRepo, err := mapsetrepository.New(cfg, lg)
	beatmapRepo, err := beatmaprepository.New(cfg, lg)
	followingRepo, err := trackingrepository.New(cfg, lg)

	worker := trackingworkerapi.New(cfg, lg, track.New(txm, userRepo, mapsetRepo, beatmapRepo, followingRepo))
	closer.Add(func() {
		worker.Start(ctx)
	})

	return nil
}
