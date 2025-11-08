package enrichesrepository

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"osu-dashboard/internal/config"
	"osu-dashboard/internal/database/repository/model"
	"osu-dashboard/internal/database/txmanager"
)

const enrichesTableName = "enriches"

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

func (r *GormRepository) Create(ctx context.Context, tx txmanager.Tx, enrich *model.Enrich) error {
	enrich.ID = 0

	err := tx.DB().WithContext(ctx).Table(enrichesTableName).Create(enrich).Error
	if err != nil {
		return fmt.Errorf("failed to create enrich: %w", err)
	}

	return nil
}

func (r *GormRepository) GetLastEnrich(ctx context.Context, tx txmanager.Tx) (*model.Enrich, error) {
	var enrich model.Enrich
	err := tx.DB().WithContext(ctx).Table(enrichesTableName).Order("enriched_at desc").Limit(1).Find(&enrich).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get last enrich: %w", err)
	}

	return &enrich, nil
}
