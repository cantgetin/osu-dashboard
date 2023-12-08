package beatmapserviceapi

import (
	"context"
	"log"
	"playcount-monitor-backend/internal/repository/model"
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
