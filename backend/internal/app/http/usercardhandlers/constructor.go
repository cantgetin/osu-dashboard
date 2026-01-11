package usercardhandlers

import (
	"context"
	"osu-dashboard/internal/dto"
	"osu-dashboard/internal/usecase/command"

	log "github.com/sirupsen/logrus"
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

func New(lg *log.Logger, c userCardCreator, p userCardProvider, u userCardUpdater) *Handlers {
	return &Handlers{
		lg:               lg,
		userCardCreator:  c,
		userCardProvider: p,
		userCardUpdater:  u,
	}
}
