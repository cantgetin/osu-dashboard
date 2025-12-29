package usercardhandlers

import (
	"context"
	log "github.com/sirupsen/logrus"
	"osu-dashboard/internal/dto"
	"osu-dashboard/internal/usecase/command"
)

type (
	userCardCreator interface {
		Create(ctx context.Context, cmd *command.CreateUserCardCommand) error
	}

	userCardProvider interface {
		Get(ctx context.Context, id int, page int) (*dto.UserCard, error)
	}

	userCardUpdater interface {
		Update(ctx context.Context, cmd *command.UpdateUserCardCommand) error
	}

	Handlers struct {
		lg               *log.Logger
		userCardCreator  userCardCreator
		userCardProvider userCardProvider
		userCardUpdater  userCardUpdater
	}
)

func New(
	lg *log.Logger,
	userCreator userCardCreator,
	userProvider userCardProvider,
	userUpdater userCardUpdater,
) *Handlers {
	return &Handlers{
		lg:               lg,
		userCardCreator:  userCreator,
		userCardProvider: userProvider,
		userCardUpdater:  userUpdater,
	}
}
