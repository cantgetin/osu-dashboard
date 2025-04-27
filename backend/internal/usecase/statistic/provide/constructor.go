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
	ListForMapsets(ctx context.Context, tx txmanager.Tx, mapsetIDs ...int) ([]*model.Beatmap, error)
	TotalCount(ctx context.Context, tx txmanager.Tx) (int, error)
}

type mapsetStore interface {
	ListForUser(ctx context.Context, tx txmanager.Tx, userId int) ([]*model.Mapset, error)
	TotalCount(ctx context.Context, tx txmanager.Tx) (int, error)
}

type userStore interface {
	TotalCount(ctx context.Context, tx txmanager.Tx) (int, error)
}

type trackStore interface {
	TotalCount(ctx context.Context, tx txmanager.Tx) (int, error)
}

type UseCase struct {
	cfg     *config.Config
	lg      *log.Logger
	txm     txmanager.TxManager
	beatmap beatmapStore
	mapset  mapsetStore
	user    userStore
	track   trackStore
}

func New(
	cfg *config.Config,
	lg *log.Logger,
	txm txmanager.TxManager,
	beatmap beatmapStore,
	mapset mapsetStore,
	user userStore,
	track trackStore,
) *UseCase {
	return &UseCase{
		cfg:     cfg,
		lg:      lg,
		txm:     txm,
		beatmap: beatmap,
		mapset:  mapset,
		track:   track,
		user:    user,
	}
}
