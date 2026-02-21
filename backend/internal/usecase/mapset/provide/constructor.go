package mapsetprovide

import (
	"context"
	"osu-dashboard/internal/config"
	"osu-dashboard/internal/database/model"
	"osu-dashboard/internal/database/txmanager"

	log "github.com/sirupsen/logrus"
)

type (
	beatmapStore interface {
		Create(ctx context.Context, tx txmanager.Tx, beatmap *model.Beatmap) error
		Update(ctx context.Context, tx txmanager.Tx, beatmap *model.Beatmap) error
		Get(ctx context.Context, tx txmanager.Tx, id int) (*model.Beatmap, error)
		ListForMapset(ctx context.Context, tx txmanager.Tx, mapsetId int) ([]*model.Beatmap, error)
		ListForMapsets(ctx context.Context, tx txmanager.Tx, mapsetIDs ...int) ([]*model.Beatmap, error)
	}

	mapsetStore interface {
		Create(ctx context.Context, tx txmanager.Tx, mapset *model.Mapset) error
		Get(ctx context.Context, tx txmanager.Tx, id int) (*model.Mapset, error)
		Update(ctx context.Context, tx txmanager.Tx, mapset *model.Mapset) error
		Exists(ctx context.Context, tx txmanager.Tx, id int) (bool, error)
		ListWithFilterSortLimitOffset(
			ctx context.Context,
			tx txmanager.Tx,
			filter model.MapsetFilter,
			sort model.MapsetSort,
			limit int,
			offset int,
		) ([]*model.Mapset, int, error)
		ListForUser(ctx context.Context, tx txmanager.Tx, userId int) ([]*model.Mapset, error)
		ListForUserWithFilterSortLimitOffset(
			ctx context.Context,
			tx txmanager.Tx,
			userID int,
			filter model.MapsetFilter,
			sort model.MapsetSort,
			limit int,
			offset int,
		) ([]*model.Mapset, int, error)
	}
)

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
