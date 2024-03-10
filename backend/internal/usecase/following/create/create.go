package followingcreate

import (
	"context"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
	"time"
)

func (uc *UseCase) Create(
	ctx context.Context,
	id int,
	username string,
) error {
	txErr := uc.txm.ReadWrite(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		follow := &model.Following{
			ID:        id,
			Username:  username,
			CreatedAt: time.Now().UTC(),
		}

		err := uc.following.Create(ctx, tx, follow)
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
