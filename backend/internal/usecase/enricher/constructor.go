package enricherusecase

import (
	"context"
	log "github.com/sirupsen/logrus"
	"osu-dashboard/internal/config"
	"osu-dashboard/internal/database/repository/model"
	"osu-dashboard/internal/database/txmanager"
	"osu-dashboard/internal/service/osuapi"
)

type mapsetStore interface {
	ListForUser(ctx context.Context, tx txmanager.Tx, userID int) ([]*model.Mapset, error)
	UpdateGenreLanguage(ctx context.Context, tx txmanager.Tx, id int, newGenre string, newLanguage string) error
}

type followingStore interface {
	List(ctx context.Context, tx txmanager.Tx) ([]*model.Following, error)
}

type enrichStore interface {
	Create(ctx context.Context, tx txmanager.Tx, enrich *model.Enrich) error
	GetLastEnrich(ctx context.Context, tx txmanager.Tx) (*model.Enrich, error)
}

type UseCase struct {
	cfg       *config.Config
	lg        *log.Logger
	txm       txmanager.TxManager
	mapset    mapsetStore
	osuApi    *osuapi.Service
	following followingStore
	enriches  enrichStore
}

func New(
	cfg *config.Config,
	lg *log.Logger,
	txManager txmanager.TxManager,
	osuAPI *osuapi.Service,
	mapset mapsetStore,
	following followingStore,
	enrichesStore enrichStore,
) *UseCase {
	return &UseCase{
		cfg:       cfg,
		lg:        lg,
		txm:       txManager,
		mapset:    mapset,
		osuApi:    osuAPI,
		following: following,
		enriches:  enrichesStore,
	}
}
