package userrepository

import (
	"context"
	"playcount-monitor-backend/internal/database/repository/model"
)

type Interface interface {
	Create(ctx context.Context, user *model.User) error
	Get(ctx context.Context, id string) (*model.User, error)
	GetByName(ctx context.Context, name string) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	List(ctx context.Context) ([]*model.User, error)
}
