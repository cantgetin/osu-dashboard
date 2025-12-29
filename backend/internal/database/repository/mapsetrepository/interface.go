package mapsetrepository

import (
	"context"
	"osu-dashboard/internal/database/repository/model"
	"osu-dashboard/internal/database/txmanager"
)

type Interface interface {
	Create(ctx context.Context, tx txmanager.Tx, mapset *model.Mapset) error
	Get(ctx context.Context, tx txmanager.Tx, id int) (*model.Mapset, error)
	Update(ctx context.Context, tx txmanager.Tx, mapset *model.Mapset) error
	Exists(ctx context.Context, tx txmanager.Tx, id int) (bool, error)
	List(ctx context.Context, tx txmanager.Tx) ([]*model.Mapset, error)
	ListForUser(ctx context.Context, tx txmanager.Tx, userID int) ([]*model.Mapset, error)
	ListForUserWithLimitOffset(ctx context.Context, tx txmanager.Tx, userID int, limit int, offset int) ([]*model.Mapset, error)
	ListStatusesForUser(ctx context.Context, tx txmanager.Tx, userID int) ([]string, error)
	ListWithFilterSortLimitOffset(
		ctx context.Context,
		tx txmanager.Tx,
		filter model.MapsetFilter,
		sort model.MapsetSort,
		limit int,
		offset int,
	) ([]*model.Mapset, int, error)
	ListForUserWithFilterSortLimitOffset(
		ctx context.Context,
		tx txmanager.Tx,
		userID int,
		filter model.MapsetFilter,
		sort model.MapsetSort,
		limit int,
		offset int,
	) ([]*model.Mapset, int, error)
	TotalCount(ctx context.Context, tx txmanager.Tx) (int, error)
	UpdateGenreLanguage(ctx context.Context, tx txmanager.Tx, id int, newGenre string, newLanguage string) error
}
