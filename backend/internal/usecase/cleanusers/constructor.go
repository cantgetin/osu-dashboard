package cleanusers

import (
	"context"
	log "github.com/sirupsen/logrus"
	"osu-dashboard/internal/config"
	"osu-dashboard/internal/database/model"
	"osu-dashboard/internal/database/txmanager"
)

type (
	userStore interface {
		List(ctx context.Context, tx txmanager.Tx) ([]*model.User, error)
		Update(ctx context.Context, tx txmanager.Tx, user *model.User) error
		Delete(ctx context.Context, tx txmanager.Tx, id int) error
	}

	mapsetStore interface {
		ListForUser(ctx context.Context, tx txmanager.Tx, userId int) ([]*model.Mapset, error)
	}

	logSource interface {
		Create(ctx context.Context, tx txmanager.Tx, log *model.Log) error
	}
)

func New(
	cfg *config.Config,
	lg *log.Logger,
	txm txmanager.TxManager,
	user userStore,
	mapset mapsetStore,
	log logSource,
) *UseCase {
	return &UseCase{
		cfg:    cfg,
		lg:     lg,
		txm:    txm,
		log:    log,
		user:   user,
		mapset: mapset,
	}
}

type UseCase struct {
	cfg    *config.Config
	lg     *log.Logger
	txm    txmanager.TxManager
	user   userStore
	mapset mapsetStore
	log    logSource
}
