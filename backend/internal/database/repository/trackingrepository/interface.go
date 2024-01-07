package trackingrepository

import (
	"context"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
)

type Interface interface {
	Create(ctx context.Context, tx txmanager.Tx, user *model.Following) error
	List(ctx context.Context, tx txmanager.Tx) ([]*model.Following, error)
	Delete(ctx context.Context, tx txmanager.Tx, id int) error
}
