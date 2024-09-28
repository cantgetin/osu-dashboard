package cleaner

import (
	"context"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
	"time"
)

func (uc *UseCase) CreateCleanRecord(
	ctx context.Context,
) error {
	track := &model.Clean{
		CleanedAt: time.Now().UTC(),
	}

	if err := uc.txm.ReadWrite(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		err := uc.clean.Create(ctx, tx, track)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}
