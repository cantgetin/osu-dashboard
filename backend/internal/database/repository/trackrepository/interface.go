package trackrepository

import (
	"context"
	"osu-dashboard/internal/database/repository/model"
	"osu-dashboard/internal/database/txmanager"
)

type Interface interface {
	Create(ctx context.Context, tx txmanager.Tx, track *model.Track) error
	GetLastTrack(ctx context.Context, tx txmanager.Tx) (*model.Track, error)
	List(ctx context.Context, tx txmanager.Tx) ([]*model.Track, error)
	TotalCount(ctx context.Context, tx txmanager.Tx) (int, error)
}
