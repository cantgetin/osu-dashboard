package followingcreate

import (
	"context"
	"playcount-monitor-backend/internal/config"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"

	log "github.com/sirupsen/logrus"
)

type followingStore interface {
	Create(ctx context.Context, tx txmanager.Tx, user *model.Following) error
}

type trackUseCase interface {
	TrackSingleFollowing(ctx context.Context, following *model.Following) error
}

type UseCase struct {
	cfg       *config.Config
	lg        *log.Logger
	txm       txmanager.TxManager
	following followingStore
	track     trackUseCase
}

func New(
	cfg *config.Config,
	lg *log.Logger,
	txm txmanager.TxManager,
	following followingStore,
	track trackUseCase,
) *UseCase {
	return &UseCase{
		cfg:       cfg,
		lg:        lg,
		txm:       txm,
		following: following,
		track:     track,
	}
}
