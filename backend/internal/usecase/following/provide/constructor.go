package followingprovide

import (
	"context"
	"osu-dashboard/internal/config"
	"osu-dashboard/internal/database/repository/model"
	"osu-dashboard/internal/database/txmanager"

	log "github.com/sirupsen/logrus"
)

type (
	followingStore interface {
		List(ctx context.Context, tx txmanager.Tx) ([]*model.Following, error)
	}

	UseCase struct {
		cfg       *config.Config
		lg        *log.Logger
		txm       txmanager.TxManager
		following followingStore
	}
)

func New(
	cfg *config.Config,
	lg *log.Logger,
	txm txmanager.TxManager,
	following followingStore,
) *UseCase {
	return &UseCase{
		cfg:       cfg,
		lg:        lg,
		txm:       txm,
		following: following,
	}
}
