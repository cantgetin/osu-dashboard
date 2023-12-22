package usercardcreate

import (
	"context"
	"playcount-monitor-backend/internal/config"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"

	log "github.com/sirupsen/logrus"
)

type userStore interface {
	Create(ctx context.Context, tx txmanager.Tx, user *model.User) error
}

type mapsetStore interface {
	Create(ctx context.Context, tx txmanager.Tx, mapset *model.Mapset) error
}

type beatmapStore interface {
	Create(ctx context.Context, tx txmanager.Tx, beatmap *model.Beatmap) error
}

type UseCase struct {
	cfg     *config.Config
	lg      *log.Logger
	txm     txmanager.TxManager
	user    userStore
	mapset  mapsetStore
	beatmap beatmapStore
}

func New(
	cfg *config.Config,
	lg *log.Logger,
	txm txmanager.TxManager,
	user userStore,
	mapset mapsetStore,
	beatmap beatmapStore,
) *UseCase {
	return &UseCase{
		cfg:     cfg,
		lg:      lg,
		txm:     txm,
		user:    user,
		mapset:  mapset,
		beatmap: beatmap,
	}
}
