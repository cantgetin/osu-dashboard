package logprovide

import (
	"context"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
	"playcount-monitor-backend/internal/dto"
)

const logsPerPage = 50

func (uc *UseCase) List(ctx context.Context, page int) (*dto.LogsPaged, error) {
	var logs []*model.Log
	var count int

	txErr := uc.txm.ReadOnly(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		var err error
		logs, count, err = uc.log.ListLogsWithLimitOffset(
			ctx,
			tx,
			logsPerPage,
			(page-1)*logsPerPage,
		)
		if err != nil {
			return err
		}

		return nil
	})
	if txErr != nil {
		return nil, txErr
	}

	return &dto.LogsPaged{
		Logs:        mapLogModelsToLogsDTOs(logs),
		CurrentPage: page,
		Pages:       (count + logsPerPage - 1) / logsPerPage,
	}, nil
}
