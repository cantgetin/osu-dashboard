package trackingcreate

import (
	"context"
	"playcount-monitor-backend/internal/config"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"

	log "github.com/sirupsen/logrus"
)

type trackingStore interface {
	Create(ctx context.Context, tx txmanager.Tx, user *model.Tracking) error
}

type UseCase struct {
	cfg      *config.Config
	lg       *log.Logger
	txm      txmanager.TxManager
	tracking trackingStore
}

func New(
	cfg *config.Config,
	lg *log.Logger,
	txm txmanager.TxManager,
	tracking trackingStore,
) *UseCase {
	return &UseCase{
		cfg:      cfg,
		lg:       lg,
		txm:      txm,
		tracking: tracking,
	}
}
