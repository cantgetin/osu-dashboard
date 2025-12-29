package loghandlers

import (
	"context"
	"osu-dashboard/internal/dto"

	log "github.com/sirupsen/logrus"
)

type (
	logProvider interface {
		List(ctx context.Context, page int) (*dto.LogsPaged, error)
	}

	Handlers struct {
		lg          *log.Logger
		logProvider logProvider
	}
)

func New(lg *log.Logger, logProvider logProvider) *Handlers {
	return &Handlers{
		lg:          lg,
		logProvider: logProvider,
	}
}
