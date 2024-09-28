package cleaner

import (
	"context"
	log "github.com/sirupsen/logrus"
	"playcount-monitor-backend/internal/config"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
)

func New(
	cfg *config.Config,
	lg *log.Logger,
	txm txmanager.TxManager,
	user userStore,
	mapset mapsetStore,
	beatmap beatmapStore,
	clean cleanStore,
) *UseCase {
	return &UseCase{
		cfg:     cfg,
		lg:      lg,
		txm:     txm,
		user:    user,
		mapset:  mapset,
		beatmap: beatmap,
		clean:   clean,
	}
}

type UseCase struct {
	cfg     *config.Config
	lg      *log.Logger
	txm     txmanager.TxManager
	user    userStore
	mapset  mapsetStore
	beatmap beatmapStore
	clean   cleanStore
}

type userStore interface {
	List(ctx context.Context, tx txmanager.Tx) ([]*model.User, error)
	Update(ctx context.Context, tx txmanager.Tx, user *model.User) error
}
type mapsetStore interface {
	List(ctx context.Context, tx txmanager.Tx) ([]*model.Mapset, error)
	Update(ctx context.Context, tx txmanager.Tx, mapset *model.Mapset) error
}
type beatmapStore interface {
	ListForMapset(ctx context.Context, tx txmanager.Tx, mapsetID int) ([]*model.Beatmap, error)
	Update(ctx context.Context, tx txmanager.Tx, beatmap *model.Beatmap) error
}

type cleanStore interface {
	Create(ctx context.Context, tx txmanager.Tx, clean *model.Clean) error
	GetLastClean(ctx context.Context, tx txmanager.Tx) (*model.Clean, error)
}
