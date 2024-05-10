package app

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"playcount-monitor-backend/internal/app/trackingworkerapi"
	"playcount-monitor-backend/internal/bootstrap"
	"playcount-monitor-backend/internal/config"
	"playcount-monitor-backend/internal/database/repository/beatmaprepository"
	"playcount-monitor-backend/internal/database/repository/followingrepository"
	"playcount-monitor-backend/internal/database/repository/mapsetrepository"
	"playcount-monitor-backend/internal/database/repository/trackrepository"
	"playcount-monitor-backend/internal/database/repository/userrepository"
	"playcount-monitor-backend/internal/service/osuapi"
	"playcount-monitor-backend/internal/service/osuapitokenprovider"
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
	//closer.Add(cancel)

	db, err := bootstrap.InitDB(cfg)
	if err != nil {
		return fmt.Errorf("failed to init db: %w", err)
	}

	const waitForConnection = 5 * time.Second
	txm := bootstrap.ConnectTxManager("tracking-worker", waitForConnection, db, lg)

	// init repos
	userRepo := userrepository.New(cfg, lg)
	mapsetRepo := mapsetrepository.New(cfg, lg)
	beatmapRepo := beatmaprepository.New(cfg, lg)
	followingRepo := followingrepository.New(cfg, lg)
	trackRepo := trackrepository.New(cfg, lg)

	// init api
	httpClient := bootstrap.NewHTTPClient()
	osuTokenProvider := osuapitokenprovider.New(cfg, &httpClient)
	osuAPI := osuapi.New(cfg, osuTokenProvider, &httpClient)

	worker := trackingworkerapi.New(cfg, lg, track.New(
		cfg,
		txm,
		osuAPI,
		userRepo,
		mapsetRepo,
		beatmapRepo,
		followingRepo,
		trackRepo,
	))

	worker.Start(ctx)

	gracefulShutDown(ctx, cancel)

	return nil
}
