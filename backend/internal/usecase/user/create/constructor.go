package usercreate

import (
	"context"
	"osu-dashboard/internal/config"
	"osu-dashboard/internal/database/repository/model"
	"osu-dashboard/internal/database/txmanager"

	log "github.com/sirupsen/logrus"
)

type userStore interface {
	Create(ctx context.Context, tx txmanager.Tx, user *model.User) error
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
