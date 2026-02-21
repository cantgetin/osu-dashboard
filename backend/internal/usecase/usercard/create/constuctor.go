package usercardcreate

import (
	"context"
	"osu-dashboard/internal/config"
	"osu-dashboard/internal/database/model"
	"osu-dashboard/internal/database/txmanager"

	log "github.com/sirupsen/logrus"
)

type (
	userStore interface {
		Create(ctx context.Context, tx txmanager.Tx, user *model.User) error
	}

	mapsetStore interface {
		Create(ctx context.Context, tx txmanager.Tx, mapset *model.Mapset) error
	}

	beatmapStore interface {
		Create(ctx context.Context, tx txmanager.Tx, beatmap *model.Beatmap) error
	}

	UseCase struct {
		cfg     *config.Config
		lg      *log.Logger
		txm     txmanager.TxManager
		user    userStore
		mapset  mapsetStore
		beatmap beatmapStore
	}
)

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
