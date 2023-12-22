package userserviceapi

import (
	"context"
	log "github.com/sirupsen/logrus"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/dto"
)

type userCreator interface {
	Create(ctx context.Context, user *dto.User) error
}

type userProvider interface {
	Get(ctx context.Context, id int) (*model.User, error)
	GetByName(ctx context.Context, name string) (*model.User, error)
	List(ctx context.Context) ([]*model.User, error)
}

type userUpdater interface {
	Update(ctx context.Context, user *model.User) error
}

type ServiceImpl struct {
	lg           *log.Logger
	userProvider userProvider
	userCreator  userCreator
	userUpdater  userUpdater
}

func New(
	lg *log.Logger,
	userProvider userProvider,
	userCreator userCreator,
	userUpdater userUpdater,
) *ServiceImpl {
	return &ServiceImpl{
		lg:           lg,
		userProvider: userProvider,
		userCreator:  userCreator,
		userUpdater:  userUpdater,
	}
}
