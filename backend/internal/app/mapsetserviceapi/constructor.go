package mapsetserviceapi

import (
	"context"
	log "github.com/sirupsen/logrus"
	"playcount-monitor-backend/internal/database/repository/model"
)

type mapsetProvider interface {
	Create(ctx context.Context, cmd *model.Mapset) error
}

type ServiceImpl struct {
	lg             *log.Logger
	mapsetProvider mapsetProvider
}

func New(
	lg *log.Logger,
	mapsetProvider mapsetProvider,
) *ServiceImpl {
	return &ServiceImpl{
		mapsetProvider: mapsetProvider,
		lg:             lg,
	}
}
