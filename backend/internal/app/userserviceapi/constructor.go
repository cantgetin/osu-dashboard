package userserviceapi

import (
	"context"
	log "github.com/sirupsen/logrus"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/dto"
	userprovide "playcount-monitor-backend/internal/usecase/user/provide"
)

type userCreator interface {
	Create(ctx context.Context, user *dto.User) error
}

type userProvider interface {
	Get(ctx context.Context, id int) (*dto.User, error)
	GetByName(ctx context.Context, name string) (*dto.User, error)
	List(ctx context.Context, cmd *userprovide.ListIn) (*userprovide.ListOut, error)
}

type userUpdater interface {
	Update(ctx context.Context, user *model.User) error
}

type ServiceImpl struct {
	lg           *log.Logger
	userCreator  userCreator
	userProvider userProvider
	userUpdater  userUpdater
}

func New(
	lg *log.Logger,
	userCreator userCreator,
	userProvider userProvider,
	userUpdater userUpdater,
) *ServiceImpl {
	return &ServiceImpl{
		lg:           lg,
		userCreator:  userCreator,
		userProvider: userProvider,
		userUpdater:  userUpdater,
	}
}
