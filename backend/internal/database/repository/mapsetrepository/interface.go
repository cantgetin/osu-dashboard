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
	ListForUser(ctx context.Context, tx txmanager.Tx, userID int) ([]*model.Mapset, error)
	ListForUserWithLimitOffset(ctx context.Context, tx txmanager.Tx, userID int, limit int, offset int) ([]*model.Mapset, error)
	ListStatusesForUser(ctx context.Context, tx txmanager.Tx, userID int) ([]string, error)
}
