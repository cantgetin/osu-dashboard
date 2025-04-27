package statisticserviceapi

import (
	"context"
	log "github.com/sirupsen/logrus"
	"playcount-monitor-backend/internal/dto"
)

type ServiceImpl struct {
	lg                *log.Logger
	statisticProvider statisticProvider
}

type statisticProvider interface {
	GetForUser(ctx context.Context, id int) (*dto.UserMapStatistics, error)
	GetForSystem(ctx context.Context) (*dto.SystemStatistics, error)
}

func New(
	lg *log.Logger,
	statisticProvider statisticProvider,
) *ServiceImpl {
	return &ServiceImpl{
		lg:                lg,
		statisticProvider: statisticProvider,
	}
}
