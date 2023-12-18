package app

import (
	"context"
	"fmt"
	"github.com/ds248a/closer"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"playcount-monitor-backend/internal/app/userserviceapi"
	"playcount-monitor-backend/internal/bootstrap"
	"playcount-monitor-backend/internal/config"
	"playcount-monitor-backend/internal/database/repository/userrepository"
	"playcount-monitor-backend/internal/http"
	"syscall"
)

func Run(cfg *config.Config, lg *log.Logger) error {

	db, err := bootstrap.InitDB(cfg)
	if err != nil {
		return err
	}

	userRepo, err := userrepository.New(cfg, lg, db)
	user := userserviceapi.New(lg, userRepo)

	httpServer, err := http.New(cfg, lg, user)
	if err != nil {
		return err
	}

	_, cancel := context.WithCancel(context.Background())

	httpServer.Start()

	gracefulShutDown(cancel, lg)

	return nil
}

func gracefulShutDown(cancel context.CancelFunc, lg *log.Logger) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(ch)

	sig := <-ch
	errorMessage := fmt.Sprintf(
		"%s %v - %s",
		"Received shutdown signal:",
		sig,
		"Graceful shutdown done",
	)
	lg.Printf(errorMessage)
	cancel()
	closer.Reset()
}
