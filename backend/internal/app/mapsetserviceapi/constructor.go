package mapsetserviceapi

import (
	"context"
	log "github.com/sirupsen/logrus"
	"playcount-monitor-backend/internal/dto"
	"playcount-monitor-backend/internal/usecase/command"
)

type mapsetCreator interface {
	Create(ctx context.Context, cmd *command.CreateMapsetCommand) error
}

type mapsetProvider interface {
	Get(ctx context.Context, id int) (*dto.Mapset, error)
}

type ServiceImpl struct {
	lg             *log.Logger
	mapsetProvider mapsetProvider
	mapsetCreator  mapsetCreator
}

func New(
	lg *log.Logger,
	mapsetProvider mapsetProvider,
	mapsetCreator mapsetCreator,
) *ServiceImpl {
	return &ServiceImpl{
		mapsetCreator:  mapsetCreator,
		mapsetProvider: mapsetProvider,
		lg:             lg,
	}
}
