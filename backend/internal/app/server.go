package app

import (
	"context"
	"os"
	"os/signal"
	"playcount-monitor-backend/internal/bootstrap"
	"playcount-monitor-backend/internal/config"
	"playcount-monitor-backend/internal/database/repository/beatmaprepository"
	"playcount-monitor-backend/internal/database/repository/followingrepository"
	"playcount-monitor-backend/internal/database/repository/mapsetrepository"
	"playcount-monitor-backend/internal/database/repository/userrepository"
	"playcount-monitor-backend/internal/http"
	"playcount-monitor-backend/internal/usecase/factory"
	"syscall"
	"time"

	"github.com/ds248a/closer"
	log "github.com/sirupsen/logrus"
)

func Run(baseCtx context.Context, cfg *config.Config, lg *log.Logger) error {

	db, err := bootstrap.InitDB(cfg)
	if err != nil {
		return err
	}

	txm := bootstrap.ConnectTxManager("playcount-monitor-backend", 5, db, lg)

	// init repos
	userRepo, err := userrepository.New(cfg, lg)
	mapsetRepo, err := mapsetrepository.New(cfg, lg)
	beatmapRepo, err := beatmaprepository.New(cfg, lg)
	followingRepo, err := followingrepository.New(cfg, lg)

	// useCase factory
	f, err := factory.New(cfg, lg, txm, &factory.Repositories{
		UserRepo:      userRepo,
		BeatmapRepo:   beatmapRepo,
		MapsetRepo:    mapsetRepo,
		FollowingRepo: followingRepo,
	})
	if err != nil {
		return err
	}

	httpServer, err := http.New(cfg, lg, f)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(baseCtx)

	httpServer.Start()

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
