package mapsetrepository

import (
	"context"
	"fmt"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
)

const mapsetsTableName = "mapsets"

func (r *GormRepository) Create(ctx context.Context, tx txmanager.Tx, mapset *model.Mapset) error {
	err := tx.DB().WithContext(ctx).Table(mapsetsTableName).Create(mapset).Error
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (r *GormRepository) Get(ctx context.Context, tx txmanager.Tx, id string) (*model.Mapset, error) {
	var mapset *model.Mapset
	err := tx.DB().WithContext(ctx).Table(mapsetsTableName).Where("id = ?", id).First(&mapset).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get mapset with id %v: %w", id, err)
	}

	return mapset, nil
}

func (r *GormRepository) Update(ctx context.Context, tx txmanager.Tx, mapset *model.Mapset) error {
	err := tx.DB().WithContext(ctx).Table(mapsetsTableName).Save(mapset).Error
	if err != nil {
		return fmt.Errorf("failed to update mapset: %w", err)
	}

	return nil
}

func (r *GormRepository) Exists(ctx context.Context, tx txmanager.Tx, id string) (bool, error) {
	var count int64
	err := tx.DB().WithContext(ctx).Table(mapsetsTableName).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, fmt.Errorf("failed to check if mapset with id %v exists: %w", id, err)
	}

	return count > 0, nil
}

func (r *GormRepository) ListForUser(ctx context.Context, tx txmanager.Tx, userID string) ([]*model.Mapset, error) {
	var mapsets []*model.Mapset
	err := tx.DB().WithContext(ctx).Table(mapsetsTableName).Where("user_id = ?", userID).Find(&mapsets).Error
	if err != nil {
		return nil, fmt.Errorf("failed to list mapsets for user %v: %w", userID, err)
	}

	return mapsets, nil
}
