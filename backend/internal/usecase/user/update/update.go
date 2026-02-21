package userupdate

import (
	"context"
	"osu-dashboard/internal/database/model"
	"osu-dashboard/internal/database/txmanager"
)

func (uc *UseCase) Update(ctx context.Context, user *model.User) error {
	txErr := uc.txm.ReadWrite(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		err := uc.user.Update(ctx, tx, user)
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
