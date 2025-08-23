package logrepository

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"playcount-monitor-backend/internal/config"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
)

type GormRepository struct {
	lg  *log.Logger
	cfg *config.Config
}

func New(cfg *config.Config, lg *log.Logger) *GormRepository {
	return &GormRepository{
		lg:  lg,
		cfg: cfg,
	}
}

type Interface interface {
	Create(ctx context.Context, tx txmanager.Tx, track *model.Log) error
	List(ctx context.Context, tx txmanager.Tx) ([]*model.Log, error)
}

const trackExtendedTableName = "track_extended"

func (r *GormRepository) Create(ctx context.Context, tx txmanager.Tx, track *model.Log) error {
	track.ID = 0
	err := tx.DB().WithContext(ctx).Table(trackExtendedTableName).Create(track).Error
	if err != nil {
		return fmt.Errorf("failed to create extended track: %w", err)
	}
	return nil
}

func (r *GormRepository) List(ctx context.Context, tx txmanager.Tx) ([]*model.Log, error) {
	var tracks []*model.Log
	err := tx.DB().WithContext(ctx).Table(trackExtendedTableName).Find(&tracks).Error
	if err != nil {
		return nil, fmt.Errorf("failed to list extended tracks: %w", err)
	}

	return tracks, nil
}
