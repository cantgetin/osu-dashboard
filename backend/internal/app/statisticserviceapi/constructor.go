package statisticserviceapi

import (
	"context"
	log "github.com/sirupsen/logrus"
)

type ServiceImpl struct {
	lg                *log.Logger
	statisticProvider statisticProvider
}

type statisticProvider interface {
	GetForUser(ctx context.Context, id int) UserMapStatisticsResponse
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
