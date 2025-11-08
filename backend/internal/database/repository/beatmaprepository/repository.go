package beatmaprepository

import (
	"context"
	"fmt"
	"osu-dashboard/internal/database/repository/model"
	"osu-dashboard/internal/database/txmanager"
)

const beatmapsTableName = "beatmaps"

func (r *GormRepository) Create(ctx context.Context, tx txmanager.Tx, beatmap *model.Beatmap) error {
	err := tx.DB().WithContext(ctx).Table(beatmapsTableName).Create(beatmap).Error
	if err != nil {
		return fmt.Errorf("failed to create beatmap: %w", err)
	}

	return nil
}

func (r *GormRepository) Get(ctx context.Context, tx txmanager.Tx, id int) (*model.Beatmap, error) {
	var beatmap *model.Beatmap
	err := tx.DB().WithContext(ctx).Table(beatmapsTableName).Where("id = ?", id).First(&beatmap).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get beatmap with id %v: %w", id, err)
	}

	return beatmap, nil
}

func (r *GormRepository) Update(ctx context.Context, tx txmanager.Tx, beatmap *model.Beatmap) error {
	err := tx.DB().WithContext(ctx).Table(beatmapsTableName).Save(beatmap).Error
	if err != nil {
		return fmt.Errorf("failed to update beatmap: %w", err)
	}

	return nil
}

func (r *GormRepository) ListForMapset(ctx context.Context, tx txmanager.Tx, mapsetID int) ([]*model.Beatmap, error) {
	var beatmaps []*model.Beatmap
	err := tx.DB().WithContext(ctx).Table(beatmapsTableName).Where("mapset_id = ?", mapsetID).Find(&beatmaps).Error
	if err != nil {
		return nil, fmt.Errorf("failed to list beatmaps for mapset %v: %w", mapsetID, err)
	}

	return beatmaps, nil
}

func (r *GormRepository) ListForMapsets(ctx context.Context, tx txmanager.Tx, mapsetIDs ...int) ([]*model.Beatmap, error) {
	var beatmaps []*model.Beatmap
	err := tx.DB().WithContext(ctx).Table(beatmapsTableName).Where("mapset_id IN (?)", mapsetIDs).Find(&beatmaps).Error
	if err != nil {
		return nil, fmt.Errorf("failed to list beatmaps for mapsets %v: %w", mapsetIDs, err)
	}

	return beatmaps, nil
}

func (r *GormRepository) Exists(ctx context.Context, tx txmanager.Tx, id int) (bool, error) {
	var count int64
	err := tx.DB().WithContext(ctx).Table(beatmapsTableName).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, fmt.Errorf("failed to check if beatmap with id %v exists: %w", id, err)
	}

	return count > 0, nil
}

func (r *GormRepository) TotalCount(ctx context.Context, tx txmanager.Tx) (int, error) {
	var count int64
	err := tx.DB().WithContext(ctx).Table(beatmapsTableName).Count(&count).Error
	if err != nil {
		return 0, fmt.Errorf("failed to count beatmaps: %w", err)
	}

	return int(count), nil
}
