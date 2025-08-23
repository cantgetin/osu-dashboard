package logcreate

import (
	"context"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
)

type LogSource interface {
	Create(ctx context.Context, tx txmanager.Tx, track *model.Log) error
}

type UseCase struct {
	txm txmanager.TxManager
	log LogSource
}

func New(txm txmanager.TxManager, log LogSource) *UseCase {
	return &UseCase{
		txm: txm,
		log: log,
	}
}
