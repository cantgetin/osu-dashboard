package track

import (
	"context"
	"playcount-monitor-backend/internal/config"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
	"playcount-monitor-backend/internal/service/osuapi"
	"time"
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
	ListForUser(ctx context.Context, tx txmanager.Tx, userID int) ([]*model.Mapset, error)
}

type beatmapStore interface {
	Create(ctx context.Context, tx txmanager.Tx, beatmap *model.Beatmap) error
	Get(ctx context.Context, tx txmanager.Tx, id int) (*model.Beatmap, error)
	Update(ctx context.Context, tx txmanager.Tx, beatmap *model.Beatmap) error
	Exists(ctx context.Context, tx txmanager.Tx, id int) (bool, error)
}

type followingStore interface {
	Get(ctx context.Context, tx txmanager.Tx, id int) (*model.Following, error)
	List(ctx context.Context, tx txmanager.Tx) ([]*model.Following, error)
	SetLastFetchedForUser(ctx context.Context, tx txmanager.Tx, username string, lastFetched time.Time) error
}

type trackStore interface {
	Create(ctx context.Context, tx txmanager.Tx, track *model.Track) error
	GetLastTrack(ctx context.Context, tx txmanager.Tx) (*model.Track, error)
	List(ctx context.Context, tx txmanager.Tx) ([]*model.Track, error)
}

type LogSource interface {
	Create(ctx context.Context, log *model.Log) error
}

type UseCase struct {
	cfg       *config.Config
	txm       txmanager.TxManager
	osuApi    *osuapi.Service
	user      userStore
	mapset    mapsetStore
	beatmap   beatmapStore
	following followingStore
	track     trackStore
	log       LogSource
}

func New(
	cfg *config.Config,
	txManager txmanager.TxManager,
	osuAPI *osuapi.Service,
	user userStore,
	mapset mapsetStore,
	beatmap beatmapStore,
	following followingStore,
	track trackStore,
	log LogSource,
) *UseCase {
	return &UseCase{
		cfg:       cfg,
		txm:       txManager,
		osuApi:    osuAPI,
		user:      user,
		mapset:    mapset,
		beatmap:   beatmap,
		following: following,
		track:     track,
		log:       log,
	}
}
