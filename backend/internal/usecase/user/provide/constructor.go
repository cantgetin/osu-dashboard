package userprovide

import (
	"context"
	"osu-dashboard/internal/config"
	"osu-dashboard/internal/database/repository/model"
	"osu-dashboard/internal/database/txmanager"
	"osu-dashboard/internal/service/osuapi"

	log "github.com/sirupsen/logrus"
)

type userStore interface {
	Get(ctx context.Context, tx txmanager.Tx, id int) (*model.User, error)
	GetByName(ctx context.Context, tx txmanager.Tx, name string) (*model.User, error)
	List(ctx context.Context, tx txmanager.Tx) ([]*model.User, error)
	Exists(ctx context.Context, tx txmanager.Tx, id int) (bool, error)
	ListUsersWithFilterSortLimitOffset(
		ctx context.Context,
		tx txmanager.Tx,
		filter model.UserFilter,
		sort model.UserSort,
		limit int,
		offset int,
	) ([]*model.User, int, error)
}

type UseCase struct {
	cfg    *config.Config
	lg     *log.Logger
	txm    txmanager.TxManager
	user   userStore
	osuApi *osuapi.Service
}

func New(
	cfg *config.Config,
	lg *log.Logger,
	txm txmanager.TxManager,
	user userStore,
	osuApi *osuapi.Service,
) *UseCase {
	return &UseCase{
		cfg:    cfg,
		lg:     lg,
		txm:    txm,
		user:   user,
		osuApi: osuApi,
	}
}
