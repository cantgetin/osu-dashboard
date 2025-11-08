package app

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"playcount-monitor-backend/internal/app/jobdbcleaner"
	jobenrich "playcount-monitor-backend/internal/app/jobenrichdata"
	"playcount-monitor-backend/internal/app/jobtrackingworker"
	"playcount-monitor-backend/internal/bootstrap"
	"playcount-monitor-backend/internal/config"
	"playcount-monitor-backend/internal/database/repository/beatmaprepository"
	"playcount-monitor-backend/internal/database/repository/cleanrepository"
	"playcount-monitor-backend/internal/database/repository/enrichesrepository"
	"playcount-monitor-backend/internal/database/repository/followingrepository"
	"playcount-monitor-backend/internal/database/repository/logrepository"
	"playcount-monitor-backend/internal/database/repository/mapsetrepository"
	"playcount-monitor-backend/internal/database/repository/trackrepository"
	"playcount-monitor-backend/internal/database/repository/userrepository"
	"playcount-monitor-backend/internal/service/osuapi"
	"playcount-monitor-backend/internal/service/osuapitokenprovider"
	cleanerUseCase "playcount-monitor-backend/internal/usecase/cleaner"
	enricherusecase "playcount-monitor-backend/internal/usecase/enricher"
	logcreate "playcount-monitor-backend/internal/usecase/log/create"
	"playcount-monitor-backend/internal/usecase/track"
	"time"
)

// TODO: use factory here instead?
func RunJobs(ctx context.Context, cfg *config.Config, lg *log.Logger) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	db, err := bootstrap.InitDB(cfg)
	if err != nil {
		return fmt.Errorf("failed to init db: %w", err)
	}

	err = bootstrap.ApplyMigrations(db)
	if err != nil {
		return err
	}

	const waitForConnection = 5 * time.Second
	txm := bootstrap.ConnectTxManager("tracking-worker", waitForConnection, db, lg)

	// init repos
	userRepo := userrepository.New(cfg, lg)
	mapsetRepo := mapsetrepository.New(cfg, lg)
	beatmapRepo := beatmaprepository.New(cfg, lg)
	followingRepo := followingrepository.New(cfg, lg)
	trackRepo := trackrepository.New(cfg, lg)
	logRepo := logrepository.New(cfg, lg)
	cleanerRepo := cleanrepository.New(cfg, lg)
	logUC := logcreate.New(txm, logRepo)
	enrichesRepo := enrichesrepository.New(cfg, lg)

	// init api
	httpClient := bootstrap.NewHTTPClient()
	osuTokenProvider := osuapitokenprovider.New(cfg, httpClient)
	osuAPI := osuapi.New(cfg, osuTokenProvider, httpClient)

	// init usecases
	cleanerUC := cleanerUseCase.New(cfg, lg, txm, userRepo, mapsetRepo, beatmapRepo, logUC, cleanerRepo)
	enricherUC := enricherusecase.New(cfg, lg, txm, osuAPI, mapsetRepo, followingRepo, enrichesRepo)

	// TODO make tracker usecase and pass
	// init tracker
	tracker := trackingworker.New(cfg, lg, track.New(
		cfg,
		lg,
		txm,
		osuAPI,
		userRepo,
		mapsetRepo,
		beatmapRepo,
		followingRepo,
		trackRepo,
		logcreate.New(txm, logRepo),
	))

	// init cleaner
	cleaner := dbcleaner.New(cfg, lg, cleanerUC)

	// init enricher
	enricher := jobenrich.New(cfg, lg, enricherUC)

	// start workers
	go tracker.Start(ctx)
	go cleaner.Start(ctx)
	go enricher.Start(ctx)

	gracefulShutDown(ctx, cancel)

	return nil
}
