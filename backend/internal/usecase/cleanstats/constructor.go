package cleanstats

import (
	"context"
	"osu-dashboard/internal/config"
	"osu-dashboard/internal/database/model"
	"osu-dashboard/internal/database/txmanager"

	log "github.com/sirupsen/logrus"
)

type (
	userStore interface {
		List(ctx context.Context, tx txmanager.Tx) ([]*model.User, error)
		Update(ctx context.Context, tx txmanager.Tx, user *model.User) error
	}

	mapsetStore interface {
		List(ctx context.Context, tx txmanager.Tx) ([]*model.Mapset, error)
		Update(ctx context.Context, tx txmanager.Tx, mapset *model.Mapset) error
	}

	beatmapStore interface {
		ListForMapset(ctx context.Context, tx txmanager.Tx, mapsetID int) ([]*model.Beatmap, error)
		Update(ctx context.Context, tx txmanager.Tx, beatmap *model.Beatmap) error
	}

	logSource interface {
		Create(ctx context.Context, tx txmanager.Tx, log *model.Log) error
	}
)

func New(
	cfg *config.Config,
	lg *log.Logger,
	txm txmanager.TxManager,
	user userStore,
	mapset mapsetStore,
	beatmap beatmapStore,
	log logSource,
) *UseCase {
	return &UseCase{
		cfg:     cfg,
		lg:      lg,
		txm:     txm,
		log:     log,
		user:    user,
		mapset:  mapset,
		beatmap: beatmap,
	}
}

type UseCase struct {
	cfg     *config.Config
	lg      *log.Logger
	txm     txmanager.TxManager
	user    userStore
	mapset  mapsetStore
	beatmap beatmapStore
	log     logSource
}
