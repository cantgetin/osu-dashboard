package statisticprovide

import (
	"context"
	"osu-dashboard/internal/config"
	"osu-dashboard/internal/database/model"
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
		SumLatestStats(ctx context.Context, tx txmanager.Tx) (playCount, favouriteCount, commentsCount int, err error)
	}

	userStore interface {
		TotalCount(ctx context.Context, tx txmanager.Tx) (int, error)
	}

	UseCase struct {
		cfg     *config.Config
		lg      *log.Logger
		txm     txmanager.TxManager
		beatmap beatmapStore
		mapset  mapsetStore
		user    userStore
	}
)

func New(
	cfg *config.Config,
	lg *log.Logger,
	txm txmanager.TxManager,
	beatmap beatmapStore,
	mapset mapsetStore,
	user userStore,
) *UseCase {
	return &UseCase{
		cfg:     cfg,
		lg:      lg,
		txm:     txm,
		beatmap: beatmap,
		mapset:  mapset,
		user:    user,
	}
}
