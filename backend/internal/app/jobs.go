package app

import (
	"context"
	"fmt"
	jobcleandb "osu-dashboard/internal/app/jobdbcleaner"
	jobenrich "osu-dashboard/internal/app/jobenrichdata"
	jobtrack "osu-dashboard/internal/app/jobtrackingworker"
	"osu-dashboard/internal/bootstrap"
	"osu-dashboard/internal/config"
	repositoryfactory "osu-dashboard/internal/database/repository/factory"
	"osu-dashboard/internal/service/osuapi"
	"osu-dashboard/internal/service/osuapitokenprovider"
	"osu-dashboard/internal/usecase/factory"

	log "github.com/sirupsen/logrus"
)

func RunJobs(ctx context.Context, cfg *config.Config, lg *log.Logger) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	db, err := bootstrap.InitDB(cfg)
	if err != nil {
		return fmt.Errorf("failed to init db: %w", err)
	}

	err = bootstrap.ApplyMigrations(db)
	if err != nil {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	txm := bootstrap.ConnectTxManager("osu-dashboard-jobs", db, lg)

	// init repos
	repoFactory := repositoryfactory.New(cfg, lg)
	userRepo := repoFactory.NewUserRepository()
	mapsetRepo := repoFactory.NewMapsetRepository()
	beatmapRepo := repoFactory.NewBeatmapRepository()
	followingRepo := repoFactory.NewFollowingsRepository()
	trackRepo := repoFactory.NewTrackRepository()
	logRepo := repoFactory.NewLogsRepository()
	cleanerRepo := repoFactory.NewCleansRepository()
	enrichesRepo := repoFactory.NewEnrichesRepository()

	// init api
	httpClient := bootstrap.NewHTTPClient()
	osuTokenProvider := osuapitokenprovider.New(cfg, httpClient)
	osuAPI := osuapi.New(cfg, osuTokenProvider, httpClient)

	// init usecases
	useCaseFactory := factory.New(cfg, lg, txm, osuAPI, &factory.Repositories{
		UserRepo:      userRepo,
		BeatmapRepo:   beatmapRepo,
		MapsetRepo:    mapsetRepo,
		FollowingRepo: followingRepo,
		TrackRepo:     trackRepo,
		LogRepo:       logRepo,
		CleanRepo:     cleanerRepo,
		EnrichesRepo:  enrichesRepo,
	})
	cleanerUC := useCaseFactory.MakeCleanerUseCase()
	enricherUC := useCaseFactory.MakeEnricherUseCase()
	trackUC := useCaseFactory.MakeTrackUseCase()

	// init jobs
	tracker := jobtrack.New(cfg, lg, trackUC)
	cleaner := jobcleandb.New(cfg, lg, cleanerUC)
	enricher := jobenrich.New(cfg, lg, enricherUC)

	// start parallel job workers
	go tracker.Start(ctx)
	go cleaner.Start(ctx)
	go enricher.Start(ctx)

	gracefulShutDown(ctx, cancel)
	return nil
}
