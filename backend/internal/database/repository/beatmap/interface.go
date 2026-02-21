package beatmaprepository

import (
	"context"
	"osu-dashboard/internal/database/model"
	"osu-dashboard/internal/database/txmanager"
)

type Interface interface {
	Create(ctx context.Context, tx txmanager.Tx, beatmap *model.Beatmap) error
	Update(ctx context.Context, tx txmanager.Tx, beatmap *model.Beatmap) error
	Get(ctx context.Context, tx txmanager.Tx, id int) (*model.Beatmap, error)
	ListForMapset(ctx context.Context, tx txmanager.Tx, mapsetID int) ([]*model.Beatmap, error)
	ListForMapsets(ctx context.Context, tx txmanager.Tx, mapsetIDs ...int) ([]*model.Beatmap, error)
	Exists(ctx context.Context, tx txmanager.Tx, id int) (bool, error)
	TotalCount(ctx context.Context, tx txmanager.Tx) (int, error)
}
