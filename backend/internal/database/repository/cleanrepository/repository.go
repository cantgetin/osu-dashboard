package cleanrepository

import (
	"context"
	"fmt"
	"osu-dashboard/internal/database/repository/model"
	"osu-dashboard/internal/database/txmanager"
)

const cleanTableName = "cleans"

func (r *GormRepository) Create(ctx context.Context, tx txmanager.Tx, clean *model.Clean) error {
	// get last clean id
	// todo: fix this auto increment crap
	var maxID *int
	err := tx.DB().WithContext(ctx).Table(cleanTableName).Select("max(id)").Row().Scan(&maxID)
	if err != nil {
		return fmt.Errorf("failed to create clean: %w", err)
	}
	clean.ID = *maxID + 1

	err = tx.DB().WithContext(ctx).Table(cleanTableName).Create(clean).Error
	if err != nil {
		return fmt.Errorf("failed to create clean: %w", err)
	}

	return nil
}

func (r *GormRepository) GetLastClean(ctx context.Context, tx txmanager.Tx) (*model.Clean, error) {
	var clean model.Clean
	err := tx.DB().WithContext(ctx).Table(cleanTableName).Order("cleaned_at desc").Limit(1).Find(&clean).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get last clean: %w", err)
	}

	return &clean, nil
}
