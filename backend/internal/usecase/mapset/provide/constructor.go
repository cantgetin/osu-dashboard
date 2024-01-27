package mapsetprovide

import (
	"context"
	log "github.com/sirupsen/logrus"
	"playcount-monitor-backend/internal/config"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
)

type beatmapStore interface {
	Create(ctx context.Context, tx txmanager.Tx, beatmap *model.Beatmap) error
	Update(ctx context.Context, tx txmanager.Tx, beatmap *model.Beatmap) error
	Get(ctx context.Context, tx txmanager.Tx, id int) (*model.Beatmap, error)
	ListForMapset(ctx context.Context, tx txmanager.Tx, mapsetId int) ([]*model.Beatmap, error)
}

type mapsetStore interface {
	Create(ctx context.Context, tx txmanager.Tx, mapset *model.Mapset) error
	Get(ctx context.Context, tx txmanager.Tx, id int) (*model.Mapset, error)
	Update(ctx context.Context, tx txmanager.Tx, mapset *model.Mapset) error
	Exists(ctx context.Context, tx txmanager.Tx, id int) (bool, error)
}

type UseCase struct {
	cfg     *config.Config
	lg      *log.Logger
	txm     txmanager.TxManager
	beatmap beatmapStore
	mapset  mapsetStore
}

func New(
	cfg *config.Config,
	lg *log.Logger,
	txm txmanager.TxManager,
	beatmap beatmapStore,
	mapset mapsetStore,
) *UseCase {
	return &UseCase{
		cfg:     cfg,
		lg:      lg,
		txm:     txm,
		beatmap: beatmap,
		mapset:  mapset,
	}
}
