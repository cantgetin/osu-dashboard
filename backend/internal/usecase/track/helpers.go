package track

import (
	"context"
	"osu-dashboard/internal/database/repository/model"
	"osu-dashboard/internal/database/txmanager"
	"time"
)

func (uc *UseCase) GetLastTimeTracked(ctx context.Context) (*time.Time, error) {
	t := time.Time{}
	if err := uc.txm.ReadOnly(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		track, err := uc.track.GetLastTrack(ctx, tx)
		if err != nil {
			return err
		}

		t = track.TrackedAt

		return nil
	}); err != nil {
		return nil, err
	}

	return &t, nil
}

func (uc *UseCase) CreateTrackRecord(ctx context.Context) error {
	track := &model.Track{
		TrackedAt: time.Now().UTC(),
	}

	if err := uc.txm.ReadWrite(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		err := uc.track.Create(ctx, tx, track)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}
