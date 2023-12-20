package userrepository

import (
	"context"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
)

type Interface interface {
	Create(ctx context.Context, tx txmanager.Tx, user *model.User) error
	Update(ctx context.Context, tx txmanager.Tx, user *model.User) error
	Get(ctx context.Context, tx txmanager.Tx, id string) (*model.User, error)
	GetByName(ctx context.Context, tx txmanager.Tx, name string) (*model.User, error)
	List(ctx context.Context, tx txmanager.Tx) ([]*model.User, error)
}
