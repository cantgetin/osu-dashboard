package trackrepository

import (
	"context"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
)

type Interface interface {
	Create(ctx context.Context, tx txmanager.Tx, track *model.Track) error
	GetLastTrack(ctx context.Context, tx txmanager.Tx) (*model.Track, error)
	List(ctx context.Context, tx txmanager.Tx) ([]*model.Track, error)
	TotalCount(ctx context.Context, tx txmanager.Tx) (int, error)
}
