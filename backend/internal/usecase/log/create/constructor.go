package logcreate

import (
	"context"
	"osu-dashboard/internal/database/repository/model"
	"osu-dashboard/internal/database/txmanager"
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
