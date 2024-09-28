package cleaner

import (
	"context"
	"playcount-monitor-backend/internal/database/txmanager"
	"time"
)

func (uc *UseCase) GetLastTimeCleaned(
	ctx context.Context,
) (*time.Time, error) {
	t := time.Time{}
	if err := uc.txm.ReadOnly(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		track, err := uc.clean.GetLastClean(ctx, tx)
		if err != nil {
			return err
		}

		t = track.CleanedAt

		return nil
	}); err != nil {
		return nil, err
	}

	return &t, nil
}
