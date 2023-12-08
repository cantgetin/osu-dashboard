package app

import (
	"context"
	"fmt"
	"github.com/ds248a/closer"
	"log"
	"os"
	"os/signal"
	"playcount-monitor-backend/internal/config"
	"playcount-monitor-backend/internal/http"
	"syscall"
)

func Run(cfg *config.Config, lg *log.Logger) error {
	httpServer, err := http.New(cfg, lg)
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
