package enricherusecase

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"osu-dashboard/internal/config"
	"osu-dashboard/internal/database/repository/model"
	"osu-dashboard/internal/database/txmanager"
	"osu-dashboard/internal/service/osuapi"
	"strconv"
	"time"
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

func (uc *UseCase) GetLastTimeEnriched(ctx context.Context) (*time.Time, error) {
	t := time.Time{}
	if err := uc.txm.ReadOnly(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		enrich, err := uc.enriches.GetLastEnrich(ctx, tx)
		if err != nil {
			return err
		}

		t = enrich.EnrichedAt

		return nil
	}); err != nil {
		return nil, err
	}

	return &t, nil
}

func (uc *UseCase) CreateEnrichRecord(ctx context.Context) error {
	enrich := &model.Enrich{
		EnrichedAt: time.Now().UTC(),
	}

	if err := uc.txm.ReadWrite(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		err := uc.enriches.Create(ctx, tx, enrich)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) Enrich(ctx context.Context) error {
	// getting all users IDs
	var follows []*model.Following
	if err := uc.txm.ReadOnly(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		var err error
		if follows, err = uc.following.List(ctx, tx); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	// log
	uc.lg.Infof("got following IDs from db, %v total", len(follows))
	if len(follows) == 0 {
		return fmt.Errorf("no following users present in db")
	}

	for _, f := range follows {
		if err := uc.enrichUserMapsets(ctx, f.ID); err != nil {
			return fmt.Errorf("failed to enrich user mapsets: %v", err)
		}
	}

	return nil
}

func (uc *UseCase) enrichUserMapsets(ctx context.Context, userID int) error {
	// get user mapsets
	var userMapsets []*model.Mapset
	if err := uc.txm.ReadOnly(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		var err error
		if userMapsets, err = uc.mapset.ListForUser(ctx, tx, userID); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	// fetch extended info for each mapset in case if user added/switched genre, language of mapset
	for _, mapset := range userMapsets {
		langGenreInfo, err := uc.osuApi.GetMapsetExtended(ctx, strconv.Itoa(mapset.ID))
		if err != nil {
			return fmt.Errorf("failed to get mapset extended info from api, mapset id: %v, err: %w", mapset.ID, err)
		}

		newGenre := langGenreInfo.Genre.Name
		newLanguage := langGenreInfo.Language.Name

		// update mapset genre, language inside readwrite transaction
		txErr := uc.txm.ReadWrite(ctx, func(ctx context.Context, tx txmanager.Tx) error {
			if err = uc.mapset.UpdateGenreLanguage(ctx, tx, mapset.ID, newGenre, newLanguage); err != nil {
				return fmt.Errorf("failed to update genre language for mapset: %v, err: %w", mapset.ID, err)
			}
			return nil
		})
		if txErr != nil {
			return fmt.Errorf("failed to update genre lang, transaction failed: err: %w", err)
		}

	}
	return nil
}
