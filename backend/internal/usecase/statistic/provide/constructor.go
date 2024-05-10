package statisticprovide

import (
	"context"
	log "github.com/sirupsen/logrus"
	"playcount-monitor-backend/internal/config"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
)

type beatmapStore interface {
	ListForMapset(ctx context.Context, tx txmanager.Tx, mapsetId int) ([]*model.Beatmap, error)
}

type mapsetStore interface {
	ListForUser(ctx context.Context, tx txmanager.Tx, userId int) ([]*model.Mapset, error)
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
