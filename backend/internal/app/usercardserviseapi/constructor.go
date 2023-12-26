package usercardserviseapi

import (
	"context"
	log "github.com/sirupsen/logrus"
	usercardcreate "playcount-monitor-backend/internal/usecase/usercard/create"
	usercardprovide "playcount-monitor-backend/internal/usecase/usercard/provide"
	usercardupdate "playcount-monitor-backend/internal/usecase/usercard/update"
)

type userCardCreator interface {
	Create(ctx context.Context, cmd *usercardcreate.CreateUserCardCommand) error
}

type userCardProvider interface {
	Get(ctx context.Context, id int) (*usercardprovide.UserCard, error)
}

type userCardUpdater interface {
	Update(ctx context.Context, cmd *usercardupdate.UpdateUserCardCommand) error
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
