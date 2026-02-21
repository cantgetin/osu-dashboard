package userhandlers

import (
	"context"
	"osu-dashboard/internal/database/model"
	"osu-dashboard/internal/dto"
	userprovide "osu-dashboard/internal/usecase/user/provide"

	log "github.com/sirupsen/logrus"
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

func New(lg *log.Logger, c userCreator, p userProvider, u userUpdater) *Handlers {
	return &Handlers{
		lg:           lg,
		userCreator:  c,
		userProvider: p,
		userUpdater:  u,
	}
}
