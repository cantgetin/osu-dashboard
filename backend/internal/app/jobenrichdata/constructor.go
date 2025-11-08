package enricher

import (
	"context"
	log "github.com/sirupsen/logrus"
	"playcount-monitor-backend/internal/config"
	"time"
)

type (
	enricher interface {
		Enrich(ctx context.Context) error
		GetLastTimeEnriched(ctx context.Context) (*time.Time, error)
		CreateEnrichRecord(ctx context.Context) error
	}

	Worker struct {
		cfg      *config.Config
		lg       *log.Logger
		enricher enricher
	}
)

func New(
	cfg *config.Config,
	lg *log.Logger,
	enrich enricher,
) *Worker {
	return &Worker{
		cfg:      cfg,
		lg:       lg,
		enricher: enrich,
	}
}
