package enricherusecase

import (
	"context"
	"osu-dashboard/internal/config"
	"osu-dashboard/internal/database/model"
	"osu-dashboard/internal/database/txmanager"
	"osu-dashboard/internal/service/osuapi"

	log "github.com/sirupsen/logrus"
)

type (
	mapsetStore interface {
		ListForUser(ctx context.Context, tx txmanager.Tx, userID int) ([]*model.Mapset, error)
		UpdateGenreLanguage(ctx context.Context, tx txmanager.Tx, id int, newGenre string, newLanguage string) error
	}

	followingStore interface {
		List(ctx context.Context, tx txmanager.Tx) ([]*model.Following, error)
	}

	logStore interface {
		Create(ctx context.Context, tx txmanager.Tx, log *model.Log) error
	}
)

type UseCase struct {
	cfg       *config.Config
	lg        *log.Logger
	txm       txmanager.TxManager
	mapset    mapsetStore
	osuApi    *osuapi.Service
	following followingStore
	log       logStore
}

func New(
	cfg *config.Config,
	lg *log.Logger,
	txManager txmanager.TxManager,
	osuAPI *osuapi.Service,
	mapset mapsetStore,
	following followingStore,
	logStore logStore,
) *UseCase {
	return &UseCase{
		cfg:       cfg,
		lg:        lg,
		txm:       txManager,
		mapset:    mapset,
		osuApi:    osuAPI,
		following: following,
		log:       logStore,
	}
}
