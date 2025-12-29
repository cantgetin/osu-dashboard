package logserviceapi

import (
	"context"
	log "github.com/sirupsen/logrus"
	"osu-dashboard/internal/dto"
)

type logProvider interface {
	List(ctx context.Context, page int) (*dto.LogsPaged, error)
}

type Handlers struct {
	lg          *log.Logger
	logProvider logProvider
}

func New(
	lg *log.Logger,
	logProvider logProvider,
) *Handlers {
	return &Handlers{
		lg:          lg,
		logProvider: logProvider,
	}
}
