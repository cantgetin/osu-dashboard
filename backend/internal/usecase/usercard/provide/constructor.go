package usercardprovide

import (
	"context"
	log "github.com/sirupsen/logrus"
	"osu-dashboard/internal/config"
	"osu-dashboard/internal/database/repository/model"
	"osu-dashboard/internal/database/txmanager"
)

type userStore interface {
	Get(ctx context.Context, tx txmanager.Tx, id int) (*model.User, error)
}

type mapsetStore interface {
	Get(ctx context.Context, tx txmanager.Tx, id int) (*model.Mapset, error)
	ListForUser(ctx context.Context, tx txmanager.Tx, userID int) ([]*model.Mapset, error)
	ListForUserWithLimitOffset(ctx context.Context, tx txmanager.Tx, userID int, limit int, offset int) ([]*model.Mapset, error)
	ListStatusesForUser(ctx context.Context, tx txmanager.Tx, userID int) ([]string, error)
}

type beatmapStore interface {
	Get(ctx context.Context, tx txmanager.Tx, id int) (*model.Beatmap, error)
	ListForMapset(ctx context.Context, tx txmanager.Tx, mapsetID int) ([]*model.Beatmap, error)
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
