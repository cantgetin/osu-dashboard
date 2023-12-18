package mapsetrepository

import (
	"context"
	"playcount-monitor-backend/internal/database/repository/model"
)

type Interface interface {
	Create(ctx context.Context, mapset *model.Mapset) error
	Get(ctx context.Context, id int) (*model.Mapset, error)
	Update(ctx context.Context, mapset *model.Mapset) error
	ListForUser(ctx context.Context, userID string) ([]*model.Mapset, error)
}
