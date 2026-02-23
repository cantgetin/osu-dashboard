package cleanusers

import (
	"context"
	"fmt"
	"osu-dashboard/internal/database/model"
	"osu-dashboard/internal/database/txmanager"
	"time"
)

func (uc *UseCase) Execute(ctx context.Context) error {
	started := time.Now()

	// list users
	var users []*model.User
	txErr := uc.txm.ReadWrite(ctx, func(ctx context.Context, tx txmanager.Tx) (err error) {
		users, err = uc.user.List(ctx, tx)
		if err != nil {
			return fmt.Errorf("failed to list users, err: %v", err)
		}
		return nil
	})
	if txErr != nil {
		return txErr
	}

	// delete users that have 0 mapsets
	for _, user := range users {
		txErr = uc.txm.ReadWrite(ctx, func(ctx context.Context, tx txmanager.Tx) error {
			userMapsets, err := uc.mapset.ListForUser(ctx, tx, user.ID)
			if err != nil {
				return fmt.Errorf("failed to list mapsets for user with id %v, err: %v", user.ID, err)
			}

			if len(userMapsets) != 0 {
				return nil
			}

			err = uc.user.Delete(ctx, tx, user.ID)
			if err != nil {
				return fmt.Errorf("failed to delete user with id %v, err: %v", user.ID, err)
			}

			return nil
		})
		if txErr != nil {
			return txErr
		}
	}

	// create log
	txErr = uc.txm.ReadWrite(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		if err := uc.log.Create(ctx, tx, &model.Log{
			Name:               "Daily user cleaning",
			Message:            model.LogMessageDailyUserClean,
			Service:            "db-cleaner",
			AppVersion:         "v1.0",
			Platform:           "Backend",
			Type:               model.LogTypeRegular,
			APIRequests:        0,
			SuccessRatePercent: 100,
			TrackedAt:          time.Now().UTC(),
			AvgResponseTime:    0,
			ElapsedTime:        time.Since(started),
			TimeSinceLastTrack: 0,
		}); err != nil {
			return err
		}
		return nil
	})
	if txErr != nil {
		return fmt.Errorf("failed to create clean log: %w", txErr)
	}

	return nil
}
