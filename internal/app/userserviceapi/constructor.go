package userserviceapi

import (
	"context"
	"log"
	"playcount-monitor-backend/internal/repository/model"
)

type userProvider interface {
	Create(ctx context.Context, user *model.User) error
	Get(ctx context.Context, id string) (*model.User, error)
	GetByName(ctx context.Context, name string) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	List(ctx context.Context) ([]*model.User, error)
}

type ServiceImpl struct {
	lg           *log.Logger
	userProvider userProvider
}

func New(
	lg *log.Logger,
	userProvider userProvider,
) *ServiceImpl {
	return &ServiceImpl{
		userProvider: userProvider,
		lg:           lg,
	}
}
