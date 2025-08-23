package logprovide

import (
	"context"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
)

type LogSource interface {
	List(ctx context.Context, tx txmanager.Tx) ([]*model.Log, error)
	ListLogsWithLimitOffset(ctx context.Context, tx txmanager.Tx, limit, offset int) ([]*model.Log, int, error)
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
