package userupdate

import (
	"context"
	"playcount-monitor-backend/internal/config"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"

	log "github.com/sirupsen/logrus"
)

type userStore interface {
	Update(ctx context.Context, tx txmanager.Tx, user *model.User) error
	Exists(ctx context.Context, tx txmanager.Tx, id int) (bool, error)
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
