package followingprovide

import (
	"context"
	"playcount-monitor-backend/internal/config"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"

	log "github.com/sirupsen/logrus"
)

type followingStore interface {
	List(ctx context.Context, tx txmanager.Tx) ([]*model.Following, error)
}

type UseCase struct {
	cfg       *config.Config
	lg        *log.Logger
	txm       txmanager.TxManager
	following followingStore
}

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
