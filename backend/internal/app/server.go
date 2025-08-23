package app

import (
	"context"
	"os"
	"os/signal"
	"playcount-monitor-backend/internal/bootstrap"
	"playcount-monitor-backend/internal/config"
	"playcount-monitor-backend/internal/database/repository/beatmaprepository"
	"playcount-monitor-backend/internal/database/repository/followingrepository"
	"playcount-monitor-backend/internal/database/repository/logrepository"
	"playcount-monitor-backend/internal/database/repository/mapsetrepository"
	"playcount-monitor-backend/internal/database/repository/trackrepository"
	"playcount-monitor-backend/internal/database/repository/userrepository"
	"playcount-monitor-backend/internal/http"
	"playcount-monitor-backend/internal/service/osuapi"
	"playcount-monitor-backend/internal/service/osuapitokenprovider"
	"playcount-monitor-backend/internal/usecase/factory"
	"syscall"
	"time"

	"github.com/ds248a/closer"
	log "github.com/sirupsen/logrus"
	netHttp "net/http"
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

	txm := bootstrap.ConnectTxManager("playcount-monitor-backend", 5, db, lg)

	userRepo := userrepository.New(cfg, lg)
	mapsetRepo := mapsetrepository.New(cfg, lg)
	beatmapRepo := beatmaprepository.New(cfg, lg)
	followingRepo := followingrepository.New(cfg, lg)
	trackRepo := trackrepository.New(cfg, lg)
	logRepo := logrepository.New(cfg, lg)

	// init api
	httpClient := netHttp.Client{}
	osuTokenProvider := osuapitokenprovider.New(cfg, &httpClient)
	osuAPI := osuapi.New(cfg, osuTokenProvider, &httpClient)

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
