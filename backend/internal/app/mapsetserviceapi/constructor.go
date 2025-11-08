package mapsetserviceapi

import (
	"context"
	log "github.com/sirupsen/logrus"
	"osu-dashboard/internal/dto"
	"osu-dashboard/internal/usecase/command"
	mapsetprovide "osu-dashboard/internal/usecase/mapset/provide"
)

type mapsetCreator interface {
	Create(ctx context.Context, cmd *command.CreateMapsetCommand) error
}

type mapsetProvider interface {
	Get(ctx context.Context, id int) (*dto.Mapset, error)
	List(ctx context.Context, cmd *mapsetprovide.ListCommand) (*mapsetprovide.ListResponse, error)
	ListForUser(ctx context.Context, userID int, cmd *mapsetprovide.ListCommand) (*mapsetprovide.ListResponse, error)
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
