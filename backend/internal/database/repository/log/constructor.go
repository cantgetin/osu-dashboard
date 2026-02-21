package logrepository

import (
	"context"
	"osu-dashboard/internal/config"
	"osu-dashboard/internal/database/model"
	"osu-dashboard/internal/database/txmanager"

	log "github.com/sirupsen/logrus"
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
	Create(ctx context.Context, tx txmanager.Tx, log *model.Log) error
	List(ctx context.Context, tx txmanager.Tx) ([]*model.Log, error)
	ListLogsWithLimitOffset(ctx context.Context, tx txmanager.Tx, limit, offset int) ([]*model.Log, int, error)
}
