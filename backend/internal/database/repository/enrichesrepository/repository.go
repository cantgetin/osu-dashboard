package enrichesrepository

import (
	"context"
	"fmt"
	"osu-dashboard/internal/config"
	"osu-dashboard/internal/database/repository/model"
	"osu-dashboard/internal/database/txmanager"

	log "github.com/sirupsen/logrus"
)

const enrichesTableName = "enriches"

type GormRepository struct {
	lg  *log.Logger
	cfg *config.Config
}

type Interface interface {
	Create(ctx context.Context, tx txmanager.Tx, enrich *model.Enrich) error
	GetLastEnrich(ctx context.Context, tx txmanager.Tx) (*model.Enrich, error)
}

func New(cfg *config.Config, lg *log.Logger) *GormRepository {
	return &GormRepository{
		lg:  lg,
		cfg: cfg,
	}
}

func (r *GormRepository) Create(ctx context.Context, tx txmanager.Tx, enrich *model.Enrich) error {
	// todo: fix this auto increment crap
	var maxID *int
	err := tx.DB().WithContext(ctx).Table(enrichesTableName).Select("max(id)").Row().Scan(&maxID)
	if err != nil {
		return fmt.Errorf("failed to create clean: %w", err)
	}
	enrich.ID = *maxID + 1

	err = tx.DB().WithContext(ctx).Table(enrichesTableName).Create(enrich).Error
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
