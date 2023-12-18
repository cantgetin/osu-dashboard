package beatmaprepository

import (
	"context"
	"playcount-monitor-backend/internal/database/repository/model"
)

type Interface interface {
	Create(ctx context.Context, beatmap *model.Beatmap) error
	Update(ctx context.Context, beatmap *model.Beatmap) error
	Get(ctx context.Context, id int) (*model.Beatmap, error)
	ListForMapset(ctx context.Context, mapsetID int) ([]*model.Beatmap, error)
}
