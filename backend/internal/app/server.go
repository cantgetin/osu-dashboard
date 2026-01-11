package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"osu-dashboard/internal/bootstrap"
	"osu-dashboard/internal/config"
	repositoryfactory "osu-dashboard/internal/database/repository/factory"
	"osu-dashboard/internal/http"
	"osu-dashboard/internal/service/osuapi"
	"osu-dashboard/internal/service/osuapitokenprovider"
	"osu-dashboard/internal/usecase/factory"
	"syscall"
	"time"

	"github.com/ds248a/closer"
	log "github.com/sirupsen/logrus"
)

func Run(baseCtx context.Context, cfg *config.Config, lg *log.Logger) error {
	ctx, cancel := context.WithCancel(baseCtx)
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

	// init repos
	repoFactory := repositoryfactory.New(cfg, lg)
	userRepo := repoFactory.NewUserRepository()
	mapsetRepo := repoFactory.NewMapsetRepository()
	beatmapRepo := repoFactory.NewBeatmapRepository()
	followingRepo := repoFactory.NewFollowingsRepository()
	trackRepo := repoFactory.NewTrackRepository()
	logRepo := repoFactory.NewLogsRepository()

	// init api
	httpClient := bootstrap.NewHTTPClient()
	osuTokenProvider := osuapitokenprovider.New(cfg, httpClient)
	osuAPI := osuapi.New(cfg, osuTokenProvider, httpClient)

	// usecase factory
	f := factory.New(cfg, lg, txm, osuAPI, &factory.Repositories{
		UserRepo:      userRepo,
		BeatmapRepo:   beatmapRepo,
		MapsetRepo:    mapsetRepo,
		FollowingRepo: followingRepo,
		TrackRepo:     trackRepo,
		LogRepo:       logRepo,
	})

	// setup http routes
	httpServer := http.New(cfg, lg, f)
	httpServer.Start()

	// TODO: setup grpc and gql

	gracefulShutDown(ctx, cancel)
	return nil
}

func gracefulShutDown(ctx context.Context, cancel context.CancelFunc) {
	const waitTime = 5 * time.Second

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(ch)

	select {
	case sig := <-ch:
		log.Printf("os signal received %s", sig.String())
	case <-ctx.Done():
		log.Printf("ctx done %s", ctx.Err().Error())
	}

	cancel()
	closer.Reset()
	time.Sleep(waitTime)
}
