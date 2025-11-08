package usercardserviseapi

import (
	"context"
	log "github.com/sirupsen/logrus"
	"osu-dashboard/internal/dto"
	"osu-dashboard/internal/usecase/command"
)

type userCardCreator interface {
	Create(ctx context.Context, cmd *command.CreateUserCardCommand) error
}

type userCardProvider interface {
	Get(ctx context.Context, id int, page int) (*dto.UserCard, error)
}

type userCardUpdater interface {
	Update(ctx context.Context, cmd *command.UpdateUserCardCommand) error
}

type ServiceImpl struct {
	lg               *log.Logger
	userCardCreator  userCardCreator
	userCardProvider userCardProvider
	userCardUpdater  userCardUpdater
}

func New(
	lg *log.Logger,
	userCreator userCardCreator,
	userProvider userCardProvider,
	userUpdater userCardUpdater,
) *ServiceImpl {
	return &ServiceImpl{
		lg:               lg,
		userCardCreator:  userCreator,
		userCardProvider: userProvider,
		userCardUpdater:  userUpdater,
	}
}
