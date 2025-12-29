package statistichandlers

import (
	"context"
	"osu-dashboard/internal/dto"

	log "github.com/sirupsen/logrus"
)

type (
	Handlers struct {
		lg                *log.Logger
		statisticProvider statisticProvider
	}

	statisticProvider interface {
		GetForUser(ctx context.Context, id int) (*dto.UserMapStatistics, error)
		GetForSystem(ctx context.Context) (*dto.SystemStatistics, error)
	}
)

func New(lg *log.Logger, statisticProvider statisticProvider) *Handlers {
	return &Handlers{
		lg:                lg,
		statisticProvider: statisticProvider,
	}
}
