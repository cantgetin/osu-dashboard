package app

import (
	"context"
	"os"
	"os/signal"
	"osu-dashboard/internal/bootstrap"
	"osu-dashboard/internal/config"
	"osu-dashboard/internal/database/repository/beatmaprepository"
	"osu-dashboard/internal/database/repository/followingrepository"
	"osu-dashboard/internal/database/repository/logrepository"
	"osu-dashboard/internal/database/repository/mapsetrepository"
	"osu-dashboard/internal/database/repository/trackrepository"
	"osu-dashboard/internal/database/repository/userrepository"
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
	db, err := bootstrap.InitDB(cfg)
	if err != nil {
		return err
	}

	err = bootstrap.ApplyMigrations(db)
	if err != nil {
		return err
	}

	txm := bootstrap.ConnectTxManager("osu-dashboard-bff", 5, db, lg)

	userRepo := userrepository.New(cfg, lg)
	mapsetRepo := mapsetrepository.New(cfg, lg)
	beatmapRepo := beatmaprepository.New(cfg, lg)
	followingRepo := followingrepository.New(cfg, lg)
	trackRepo := trackrepository.New(cfg, lg)
	logRepo := logrepository.New(cfg, lg)

	// init api
	httpClient := bootstrap.NewHTTPClient()
	osuTokenProvider := osuapitokenprovider.New(cfg, httpClient)
	osuAPI := osuapi.New(cfg, osuTokenProvider, httpClient)

	// useCase factory
	f, err := factory.New(cfg, lg, txm, osuAPI, &factory.Repositories{
		UserRepo:      userRepo,
		BeatmapRepo:   beatmapRepo,
		MapsetRepo:    mapsetRepo,
		FollowingRepo: followingRepo,
		TrackRepo:     trackRepo,
		LogRepo:       logRepo,
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
