package mapsethandlers

import (
	"context"
	"osu-dashboard/internal/dto"
	"osu-dashboard/internal/usecase/command"
	mapsetprovide "osu-dashboard/internal/usecase/mapset/provide"

	log "github.com/sirupsen/logrus"
)

type (
	mapsetCreator interface {
		Create(ctx context.Context, cmd *command.CreateMapsetCommand) error
	}

	mapsetProvider interface {
		Get(ctx context.Context, id int) (*dto.Mapset, error)
		List(ctx context.Context, cmd *mapsetprovide.ListCommand) (*dto.MapsetsPaged, error)
		ListForUser(ctx context.Context, userID int, cmd *mapsetprovide.ListCommand) (*dto.MapsetsPaged, error)
	}

	Handlers struct {
		lg             *log.Logger
		mapsetProvider mapsetProvider
		mapsetCreator  mapsetCreator
	}
)

func New(lg *log.Logger, p mapsetProvider, c mapsetCreator) *Handlers {
	return &Handlers{
		mapsetCreator:  c,
		mapsetProvider: p,
		lg:             lg,
	}
}
