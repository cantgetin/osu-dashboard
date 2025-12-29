package statisticprovide

import (
	"context"
	"osu-dashboard/internal/config"
	"osu-dashboard/internal/database/repository/model"
	"osu-dashboard/internal/database/txmanager"

	log "github.com/sirupsen/logrus"
)

type (
	beatmapStore interface {
		ListForMapset(ctx context.Context, tx txmanager.Tx, mapsetId int) ([]*model.Beatmap, error)
		ListForMapsets(ctx context.Context, tx txmanager.Tx, mapsetIDs ...int) ([]*model.Beatmap, error)
		TotalCount(ctx context.Context, tx txmanager.Tx) (int, error)
	}

	mapsetStore interface {
		ListForUser(ctx context.Context, tx txmanager.Tx, userId int) ([]*model.Mapset, error)
		TotalCount(ctx context.Context, tx txmanager.Tx) (int, error)
	}

	userStore interface {
		TotalCount(ctx context.Context, tx txmanager.Tx) (int, error)
	}

	trackStore interface {
		TotalCount(ctx context.Context, tx txmanager.Tx) (int, error)
	}

	UseCase struct {
		cfg     *config.Config
		lg      *log.Logger
		txm     txmanager.TxManager
		beatmap beatmapStore
		mapset  mapsetStore
		user    userStore
		track   trackStore
	}
)

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
