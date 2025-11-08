package cleanrepository

import (
	"context"
	"osu-dashboard/internal/database/repository/model"
	"osu-dashboard/internal/database/txmanager"
)

type Interface interface {
	Create(ctx context.Context, tx txmanager.Tx, clean *model.Clean) error
	GetLastClean(ctx context.Context, tx txmanager.Tx) (*model.Clean, error)
}
