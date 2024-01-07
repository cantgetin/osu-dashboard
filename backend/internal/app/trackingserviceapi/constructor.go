package trackingserviceapi

import (
	"context"
	log "github.com/sirupsen/logrus"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/dto"
)

type trackingCreator interface {
	Create(ctx context.Context, cmd *model.Following) error
}

type trackingProvider interface {
	List(ctx context.Context) ([]*dto.Tracking, error)
}

type ServiceImpl struct {
	lg               *log.Logger
	trackingCreator  trackingCreator
	trackingProvider trackingProvider
}

func New(
	lg *log.Logger,
	trackingCreator trackingCreator,
	trackingProvider trackingProvider,
) *ServiceImpl {
	return &ServiceImpl{
		lg:               lg,
		trackingCreator:  trackingCreator,
		trackingProvider: trackingProvider,
	}
}
