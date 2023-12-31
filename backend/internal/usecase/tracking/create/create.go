package trackingcreate

import (
	"context"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
)

func (uc *UseCase) Create(
	ctx context.Context,
	tracking *model.Tracking,
) error {
	txErr := uc.txm.ReadWrite(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		err := uc.tracking.Create(ctx, tx, tracking)
		if err != nil {
			return err
		}

		return nil
	})
	if txErr != nil {
		return txErr
	}

	return nil
}
