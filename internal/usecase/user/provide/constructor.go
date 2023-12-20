package userprovide

import (
	"context"
	log "github.com/sirupsen/logrus"
	"playcount-monitor-backend/internal/config"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
)

type userStore interface {
	Get(ctx context.Context, tx txmanager.Tx, id string) (*model.User, error)
	GetByName(ctx context.Context, tx txmanager.Tx, name string) (*model.User, error)
	List(ctx context.Context, tx txmanager.Tx) ([]*model.User, error)
}

type UseCase struct {
	cfg  *config.Config
	lg   *log.Logger
	txm  txmanager.TxManager
	user userStore
}

func New(
	cfg *config.Config,
	lg *log.Logger,
	txm txmanager.TxManager,
	user userStore,
) *UseCase {
	return &UseCase{
		cfg:  cfg,
		lg:   lg,
		txm:  txm,
		user: user,
	}
}
