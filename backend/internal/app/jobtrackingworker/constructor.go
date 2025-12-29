package trackingworker

import (
	"context"
	"osu-dashboard/internal/config"
	"time"

	log "github.com/sirupsen/logrus"
)

type (
	tracker interface {
		TrackAllFollowings(ctx context.Context, startTime time.Time, timeSinceLastFetch time.Duration) error
		GetLastTimeTracked(ctx context.Context) (*time.Time, error)
	}

	Worker struct {
		cfg     *config.Config
		lg      *log.Logger
		tracker tracker
	}
)

func New(cfg *config.Config, lg *log.Logger, tracker tracker) *Worker {
	return &Worker{
		cfg:     cfg,
		lg:      lg,
		tracker: tracker,
	}
}
