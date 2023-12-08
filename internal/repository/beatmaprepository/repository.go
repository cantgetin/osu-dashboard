package beatmaprepository

import (
	"context"
	"fmt"
	"playcount-monitor-backend/internal/repository/model"
)

const beatmapsTableName = "beatmaps"

func (impl *GormRepository) Create(ctx context.Context, beatmap *model.Beatmap) error {
	err := impl.db.WithContext(ctx).Table(beatmapsTableName).Create(beatmap).Error
	if err != nil {
		return fmt.Errorf("failed to create beatmap: %w", err)
	}

	return nil
}

func (impl *GormRepository) Get(ctx context.Context, id int) (*model.Beatmap, error) {
	var beatmap *model.Beatmap
	err := impl.db.WithContext(ctx).Table(beatmapsTableName).Where("id = ?", id).First(&beatmap).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get beatmap with id %v: %w", id, err)
	}

	return beatmap, nil
}

func (impl *GormRepository) Update(ctx context.Context, beatmap *model.Beatmap) error {
	err := impl.db.WithContext(ctx).Table(beatmapsTableName).Save(beatmap).Error
	if err != nil {
		return fmt.Errorf("failed to update beatmap: %w", err)
	}

	return nil
}

func (impl *GormRepository) ListForMapset(ctx context.Context, mapsetID int) ([]*model.Beatmap, error) {
	var beatmaps []*model.Beatmap
	err := impl.db.WithContext(ctx).Table(beatmapsTableName).Where("mapset_id = ?", mapsetID).Find(&beatmaps).Error
	if err != nil {
		return nil, fmt.Errorf("failed to list beatmaps for mapset %v: %w", mapsetID, err)
	}

	return beatmaps, nil
}
