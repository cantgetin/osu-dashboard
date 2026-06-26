package followingrepository

import (
	"context"
	"time"

	"osu-dashboard/internal/database/model"
	"osu-dashboard/internal/database/txmanager"
)

type Interface interface {
	Get(ctx context.Context, tx txmanager.Tx, id int) (*model.Following, error)
	Create(ctx context.Context, tx txmanager.Tx, user *model.Following) error
	List(ctx context.Context, tx txmanager.Tx) ([]*model.Following, error)
	SetLastFetchedForUser(ctx context.Context, tx txmanager.Tx, username string, lastFetched time.Time) error
	Delete(ctx context.Context, tx txmanager.Tx, id int) error
}
