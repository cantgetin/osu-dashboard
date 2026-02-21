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

func RunBFF(baseCtx context.Context, cfg *config.Config, lg *log.Logger) error {
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

	// setup http routes
	httpServer := http.New(cfg, lg, useCaseFactory)
	httpServer.Start()

	// TODO: setup gql

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
