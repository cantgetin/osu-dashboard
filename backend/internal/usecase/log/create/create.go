package logcreate

import (
	"context"
	"osu-dashboard/internal/database/model"
	"osu-dashboard/internal/database/txmanager"
)

func (uc *UseCase) Create(ctx context.Context, log *model.Log) error {
	if err := uc.txm.ReadWrite(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		err := uc.log.Create(ctx, tx, log)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}
