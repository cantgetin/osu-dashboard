package userrepository

import (
	"context"
	"osu-dashboard/internal/database/repository/model"
	"osu-dashboard/internal/database/txmanager"
)

type Interface interface {
	Create(ctx context.Context, tx txmanager.Tx, user *model.User) error
	Update(ctx context.Context, tx txmanager.Tx, user *model.User) error
	Get(ctx context.Context, tx txmanager.Tx, id int) (*model.User, error)
	GetByName(ctx context.Context, tx txmanager.Tx, name string) (*model.User, error)
	List(ctx context.Context, tx txmanager.Tx) ([]*model.User, error)
	Exists(ctx context.Context, tx txmanager.Tx, id int) (bool, error)
	TotalCount(ctx context.Context, tx txmanager.Tx) (int, error)
	ListUsersWithFilterSortLimitOffset(
		ctx context.Context,
		tx txmanager.Tx,
		filter model.UserFilter,
		sort model.UserSort,
		limit int,
		offset int,
	) ([]*model.User, int, error)
}
