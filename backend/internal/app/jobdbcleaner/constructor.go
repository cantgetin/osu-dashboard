package dbcleaner

import (
	"context"
	log "github.com/sirupsen/logrus"
	"playcount-monitor-backend/internal/config"
	"time"
)

type (
	cleaner interface {
		Clean(ctx context.Context) error
		GetLastTimeCleaned(ctx context.Context) (*time.Time, error)
		CreateCleanRecord(ctx context.Context) error
	}

	Worker struct {
		cfg     *config.Config
		lg      *log.Logger
		cleaner cleaner
	}
)

func New(
	cfg *config.Config,
	lg *log.Logger,
	cleaner cleaner,
) *Worker {
	return &Worker{
		cfg:     cfg,
		lg:      lg,
		cleaner: cleaner,
	}
}
