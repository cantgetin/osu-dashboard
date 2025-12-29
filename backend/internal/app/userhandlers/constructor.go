package userhandlers

import (
	"context"
	log "github.com/sirupsen/logrus"
	"osu-dashboard/internal/database/repository/model"
	"osu-dashboard/internal/dto"
	userprovide "osu-dashboard/internal/usecase/user/provide"
)

type (
	userCreator interface {
		Create(ctx context.Context, user *dto.User) error
	}

	userProvider interface {
		Get(ctx context.Context, id int) (*dto.User, error)
		GetByName(ctx context.Context, name string) (*dto.User, error)
		List(ctx context.Context, cmd *userprovide.ListIn) (*dto.UsersPaged, error)
	}

	userUpdater interface {
		Update(ctx context.Context, user *model.User) error
	}

	Handlers struct {
		lg           *log.Logger
		userCreator  userCreator
		userProvider userProvider
		userUpdater  userUpdater
	}
)

func New(
	lg *log.Logger,
	userCreator userCreator,
	userProvider userProvider,
	userUpdater userUpdater,
) *Handlers {
	return &Handlers{
		lg:           lg,
		userCreator:  userCreator,
		userProvider: userProvider,
		userUpdater:  userUpdater,
	}
}
