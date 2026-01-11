package searchusecase

import (
	"context"
	"osu-dashboard/internal/database/repository/model"
	"osu-dashboard/internal/database/txmanager"
)

type (
	userProvider interface {
		ListUsersWithFilterAndLimit(
			ctx context.Context,
			tx txmanager.Tx,
			filter model.UserFilter,
			limit int,
		) (users []*model.User, err error)
	}
	mapsetProvider interface {
		ListWithFilterAndLimit(
			ctx context.Context,
			tx txmanager.Tx,
			filter model.MapsetFilter,
			limit int,
		) ([]*model.Mapset, error)
	}

	UseCase struct {
		txm    txmanager.TxManager
		user   userProvider
		mapset mapsetProvider
	}
)

func New(txm txmanager.TxManager, user userProvider, mapset mapsetProvider) *UseCase {
	return &UseCase{
		txm:    txm,
		user:   user,
		mapset: mapset,
	}
}
