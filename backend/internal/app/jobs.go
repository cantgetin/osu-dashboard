package app

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"osu-dashboard/internal/app/jobdbcleaner"
	jobenrich "osu-dashboard/internal/app/jobenrichdata"
	"osu-dashboard/internal/app/jobtrackingworker"
	"osu-dashboard/internal/bootstrap"
	"osu-dashboard/internal/config"
	"osu-dashboard/internal/database/repository/beatmaprepository"
	"osu-dashboard/internal/database/repository/cleanrepository"
	"osu-dashboard/internal/database/repository/enrichesrepository"
	"osu-dashboard/internal/database/repository/followingrepository"
	"osu-dashboard/internal/database/repository/logrepository"
	"osu-dashboard/internal/database/repository/mapsetrepository"
	"osu-dashboard/internal/database/repository/trackrepository"
	"osu-dashboard/internal/database/repository/userrepository"
	"osu-dashboard/internal/service/osuapi"
	"osu-dashboard/internal/service/osuapitokenprovider"
	cleanerUseCase "osu-dashboard/internal/usecase/cleaner"
	enricherusecase "osu-dashboard/internal/usecase/enricher"
	logcreate "osu-dashboard/internal/usecase/log/create"
	"osu-dashboard/internal/usecase/track"
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
