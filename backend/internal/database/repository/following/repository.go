package followingrepository

import (
	"context"
	"fmt"
	"osu-dashboard/internal/database/model"
	"osu-dashboard/internal/database/txmanager"
	"time"
)

const followingTableName = "following"

func (r *GormRepository) Get(ctx context.Context, tx txmanager.Tx, id int) (*model.Following, error) {
	var following *model.Following
	err := tx.DB().WithContext(ctx).Table(followingTableName).Where("id = ?", id).First(&following).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get following with id %v: %w", id, err)
	}

	return following, nil
}

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

func (r *GormRepository) SetLastFetchedForUser(
	ctx context.Context,
	tx txmanager.Tx,
	username string,
	lastFetched time.Time,
) error {
	err := tx.DB().WithContext(ctx).
		Table(followingTableName).
		Where("username = ?", username).
		Update("last_fetched", lastFetched).
		Error

	if err != nil {
		return fmt.Errorf("failed to set last fetched for following %v: %w", username, err)
	}

	return nil
}
