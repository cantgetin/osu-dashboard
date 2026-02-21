package app

import (
	"context"
	"fmt"
	job "osu-dashboard/internal/app/jobs"
	"osu-dashboard/internal/bootstrap"
	"osu-dashboard/internal/config"
	repositoryfactory "osu-dashboard/internal/database/repository/factory"
	"osu-dashboard/internal/service/osuapi"
	"osu-dashboard/internal/service/osuapitokenprovider"
	"osu-dashboard/internal/usecase/factory"

	log "github.com/sirupsen/logrus"
)

const (
	JobNameCleanStats = "clean stats"
	JobNameTrackUsers = "track users"
	JobNameEnrichData = "enrich data"
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
		LogRepo:       repoFactory.NewLogsRepository(),
		JobRepo:       repoFactory.NewJobRepository(),
	})
	cleanUseCase := useCaseFactory.MakeCleanerUseCase()
	enrichUseCase := useCaseFactory.MakeEnricherUseCase()
	trackUseCase := useCaseFactory.MakeTrackUseCase()

	// init jobs
	cleanStatsJob := job.NewPeriodic(
		lg,
		JobNameCleanStats,
		cfg.CleaningInterval,
		cfg.CleaningTimeout,
		useCaseFactory.MakeCleanRecorderUseCase(),
		cleanUseCase,
	)
	trackUsersJob := job.NewPeriodic(
		lg,
		JobNameTrackUsers,
		cfg.TrackingInterval,
		cfg.TrackingTimeout,
		useCaseFactory.MakeTrackRecorderUseCase(),
		trackUseCase,
	)
	enrichDataJob := job.NewPeriodic(
		lg,
		JobNameEnrichData,
		cfg.EnrichingInterval,
		cfg.EnrichingTimeout,
		useCaseFactory.MakeEnrichRecorderUseCase(),
		enrichUseCase,
	)

	// start parallel job workers
	go cleanStatsJob.Start(ctx)
	go trackUsersJob.Start(ctx)
	go enrichDataJob.Start(ctx)

	gracefulShutDown(ctx, cancel)
	return nil
}
