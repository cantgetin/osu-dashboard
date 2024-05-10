package statisticserviceapi

import (
	"context"
	log "github.com/sirupsen/logrus"
	statisticprovide "playcount-monitor-backend/internal/usecase/statistic/provide"
)

type ServiceImpl struct {
	lg                *log.Logger
	statisticProvider statisticProvider
}

type statisticProvider interface {
	GetForUser(ctx context.Context, id int) (*statisticprovide.UserMapStatistics, error)
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
