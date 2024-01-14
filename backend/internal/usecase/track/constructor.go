package track

import (
	"context"
	"playcount-monitor-backend/internal/config"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
	"playcount-monitor-backend/internal/service/osuapi"
)

type userStore interface {
	Create(ctx context.Context, tx txmanager.Tx, user *model.User) error
	Get(ctx context.Context, tx txmanager.Tx, id int) (*model.User, error)
	Update(ctx context.Context, tx txmanager.Tx, user *model.User) error
	Exists(ctx context.Context, tx txmanager.Tx, id int) (bool, error)
}

type mapsetStore interface {
	Create(ctx context.Context, tx txmanager.Tx, mapset *model.Mapset) error
	Get(ctx context.Context, tx txmanager.Tx, id int) (*model.Mapset, error)
	Exists(ctx context.Context, tx txmanager.Tx, id int) (bool, error)
	Update(ctx context.Context, tx txmanager.Tx, mapset *model.Mapset) error
}

type beatmapStore interface {
	Create(ctx context.Context, tx txmanager.Tx, beatmap *model.Beatmap) error
	Get(ctx context.Context, tx txmanager.Tx, id int) (*model.Beatmap, error)
	Update(ctx context.Context, tx txmanager.Tx, beatmap *model.Beatmap) error
}

type followingStore interface {
	List(ctx context.Context, tx txmanager.Tx) ([]*model.Following, error)
}

type UseCase struct {
	cfg           *config.Config
	txm           txmanager.TxManager
	osuApiService osuapi.Interface
	user          userStore
	mapset        mapsetStore
	beatmap       beatmapStore
	following     followingStore
}

func New(
	cfg *config.Config,
	txManager txmanager.TxManager,
	osuAPI osuapi.Interface,
	user userStore,
	mapset mapsetStore,
	beatmap beatmapStore,
	following followingStore,
) *UseCase {
	return &UseCase{
		cfg:           cfg,
		txm:           txManager,
		osuApiService: osuAPI,
		user:          user,
		mapset:        mapset,
		beatmap:       beatmap,
		following:     following,
	}
}
