package usercardupdate

import (
	"context"
	"playcount-monitor-backend/internal/config"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"

	log "github.com/sirupsen/logrus"
)

type userStore interface {
	Get(ctx context.Context, tx txmanager.Tx, id int) (*model.User, error)
	Update(ctx context.Context, tx txmanager.Tx, user *model.User) error
}

type mapsetStore interface {
	Create(ctx context.Context, tx txmanager.Tx, mapset *model.Mapset) error
	Exists(ctx context.Context, tx txmanager.Tx, id int) (bool, error)
	Get(ctx context.Context, tx txmanager.Tx, id int) (*model.Mapset, error)
	Update(ctx context.Context, tx txmanager.Tx, mapset *model.Mapset) error
}

type beatmapStore interface {
	Create(ctx context.Context, tx txmanager.Tx, beatmap *model.Beatmap) error
	Get(ctx context.Context, tx txmanager.Tx, id int) (*model.Beatmap, error)
	Update(ctx context.Context, tx txmanager.Tx, beatmap *model.Beatmap) error
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
