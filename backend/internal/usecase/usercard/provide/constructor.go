package usercardprovide

import (
	"context"
	"osu-dashboard/internal/config"
	"osu-dashboard/internal/database/repository/model"
	"osu-dashboard/internal/database/txmanager"

	log "github.com/sirupsen/logrus"
)

type (
	userStore interface {
		Get(ctx context.Context, tx txmanager.Tx, id int) (*model.User, error)
	}

	mapsetStore interface {
		Get(ctx context.Context, tx txmanager.Tx, id int) (*model.Mapset, error)
		ListForUser(ctx context.Context, tx txmanager.Tx, userID int) ([]*model.Mapset, error)
		ListForUserWithLimitOffset(ctx context.Context, tx txmanager.Tx, userID int, limit int, offset int) ([]*model.Mapset, error)
		ListStatusesForUser(ctx context.Context, tx txmanager.Tx, userID int) ([]string, error)
	}

	beatmapStore interface {
		Get(ctx context.Context, tx txmanager.Tx, id int) (*model.Beatmap, error)
		ListForMapset(ctx context.Context, tx txmanager.Tx, mapsetID int) ([]*model.Beatmap, error)
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
