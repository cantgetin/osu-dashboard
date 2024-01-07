package trackingrepository

import (
	"context"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
)

func (g GormRepository) List(ctx context.Context, tx txmanager.Tx) ([]*model.Following, error) {
	//TODO implement me
	panic("implement me")
}
