package app

import (
	"context"
	"fmt"
	jobcleandb "osu-dashboard/internal/app/jobs/jobdbcleaner"
	jobenrich "osu-dashboard/internal/app/jobs/jobenrichdata"
	jobtrack "osu-dashboard/internal/app/jobs/jobtrackingworker"
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

	txm := bootstrap.ConnectTxManager(db, lg)

	// init api
	httpClient := bootstrap.NewHTTPClient()
	osuTokenProvider := osuapitokenprovider.New(cfg, httpClient)
	osuAPI := osuapi.New(cfg, osuTokenProvider, httpClient)

	// init repos, usecases
	repoFactory := repositoryfactory.New(cfg, lg)
	useCaseFactory := factory.New(cfg, lg, txm, osuAPI, &factory.Repositories{
		UserRepo:      repoFactory.NewUserRepository(),
		BeatmapRepo:   repoFactory.NewBeatmapRepository(),
		MapsetRepo:    repoFactory.NewMapsetRepository(),
		FollowingRepo: repoFactory.NewFollowingsRepository(),
		TrackRepo:     repoFactory.NewTrackRepository(),
		LogRepo:       repoFactory.NewLogsRepository(),
		CleanRepo:     repoFactory.NewCleansRepository(),
		EnrichesRepo:  repoFactory.NewEnrichesRepository(),
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
