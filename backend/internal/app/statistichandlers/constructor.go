package statistichandlers

import (
	"context"
	log "github.com/sirupsen/logrus"
	"osu-dashboard/internal/dto"
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

func New(
	lg *log.Logger,
	statisticProvider statisticProvider,
) *Handlers {
	return &Handlers{
		lg:                lg,
		statisticProvider: statisticProvider,
	}
}
