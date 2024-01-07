package trackingworkerapi

import (
	"context"
	log "github.com/sirupsen/logrus"
	"playcount-monitor-backend/internal/config"
)

type (
	tracker interface {
		Track(ctx context.Context) error
	}

	Worker struct {
		cfg     *config.Config
		lg      *log.Logger
		tracker tracker
	}
)

func New(
	cfg *config.Config,
	lg *log.Logger,
	tracker tracker,
) *Worker {
	return &Worker{
		cfg:     cfg,
		lg:      lg,
		tracker: tracker,
	}
}
