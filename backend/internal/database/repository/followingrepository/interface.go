package followingrepository

import (
	"context"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
	"time"
)

type Interface interface {
	Get(ctx context.Context, tx txmanager.Tx, id int) (*model.Following, error)
	Create(ctx context.Context, tx txmanager.Tx, user *model.Following) error
	List(ctx context.Context, tx txmanager.Tx) ([]*model.Following, error)
	SetLastFetchedForUser(ctx context.Context, tx txmanager.Tx, username string, lastFetched time.Time) error
	Delete(ctx context.Context, tx txmanager.Tx, id int) error
}
