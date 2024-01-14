package followingrepository

import (
	"context"
	"fmt"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
)

const followingTableName = "following"

func (r *GormRepository) Create(ctx context.Context, tx txmanager.Tx, follow *model.Following) error {
	err := tx.DB().WithContext(ctx).Table(followingTableName).Create(follow).Error
	if err != nil {
		return fmt.Errorf("failed to create follow: %w", err)
	}

	return nil
}

func (r *GormRepository) List(ctx context.Context, tx txmanager.Tx) ([]*model.Following, error) {
	var follows []*model.Following
	err := tx.DB().WithContext(ctx).Table(followingTableName).Find(&follows).Error
	if err != nil {
		return nil, fmt.Errorf("failed to list follows: %w", err)
	}

	return follows, nil
}

func (r *GormRepository) Delete(ctx context.Context, tx txmanager.Tx, id int) error {
	err := tx.DB().WithContext(ctx).Table(followingTableName).Where("id = ?", id).Delete(&model.Following{}).Error
	if err != nil {
		return fmt.Errorf("failed to delete follow with id %v: %w", id, err)
	}

	return nil
}
