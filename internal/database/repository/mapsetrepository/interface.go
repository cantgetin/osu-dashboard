package mapsetrepository

import (
	"context"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
)

type Interface interface {
	Create(ctx context.Context, tx txmanager.Tx, mapset *model.Mapset) error
	Get(ctx context.Context, tx txmanager.Tx, id int) (*model.Mapset, error)
	Update(ctx context.Context, tx txmanager.Tx, mapset *model.Mapset) error
	Exists(ctx context.Context, tx txmanager.Tx, id int) (bool, error)
	ListForUser(ctx context.Context, tx txmanager.Tx, userID string) ([]*model.Mapset, error)
}
