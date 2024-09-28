package cleanrepository

import (
	"context"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
)

type Interface interface {
	Create(ctx context.Context, tx txmanager.Tx, clean *model.Clean) error
	GetLastClean(ctx context.Context, tx txmanager.Tx) (*model.Clean, error)
}
