package trackrepository

import (
	"context"
	"fmt"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
)

const trackTableName = "tracks"

func (r *GormRepository) Create(ctx context.Context, tx txmanager.Tx, track *model.Track) error {
	err := tx.DB().WithContext(ctx).Table(trackTableName).Create(track).Error
	if err != nil {
		return fmt.Errorf("failed to create track: %w", err)
	}

	return nil
}

func (r *GormRepository) GetLastTrack(ctx context.Context, tx txmanager.Tx) (*model.Track, error) {
	var track model.Track
	err := tx.DB().WithContext(ctx).Table(trackTableName).Order("tracked_at desc").Limit(1).Find(&track).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get last track: %w", err)
	}

	return &track, nil
}

func (r *GormRepository) List(ctx context.Context, tx txmanager.Tx) ([]*model.Track, error) {
	var tracks []*model.Track
	err := tx.DB().WithContext(ctx).Table(trackTableName).Find(&tracks).Error
	if err != nil {
		return nil, fmt.Errorf("failed to list tracks: %w", err)
	}

	return tracks, nil
}
