package logrepository

import (
	"context"
	log "github.com/sirupsen/logrus"
	"osu-dashboard/internal/config"
	"osu-dashboard/internal/database/repository/model"
	"osu-dashboard/internal/database/txmanager"
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
