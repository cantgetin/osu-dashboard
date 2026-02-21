package track

import (
	"context"
	"fmt"
	"osu-dashboard/internal/database/model"
	"osu-dashboard/internal/database/txmanager"
	"strconv"
	"time"

	"golang.org/x/sync/errgroup"
)

func (uc *UseCase) Execute(ctx context.Context) error {
	startTime := time.Now()

	// get all following IDs from db and get updated data from api, update data in db
	var follows []*model.Following
	if err := uc.txm.ReadOnly(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		var err error
		if follows, err = uc.following.List(ctx, tx); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	uc.lg.Infof("got following IDs from db, %v total", len(follows))
	if len(follows) == 0 {
		return fmt.Errorf("no following users present in db")
	}

	// worker pool to limit concurrent goroutines and prevent overwhelming the API
	eg, _ := errgroup.WithContext(ctx)
	eg.SetLimit(uc.cfg.TrackingMaxParallelWorkers)

	for i, f := range follows {
		eg.Go(func() (errG error) {
			uc.lg.Infof("fetching user %s with id %v, %v/%v", f.Username, f.ID, i+1, len(follows))

			if err := uc.TrackSingleFollowing(ctx, f); err != nil {
				uc.lg.Infof("failed to fetch specific user: %s", err.Error())
			}
			return nil
		})
	}

	err := eg.Wait()
	if err != nil {
		return err
	}

	defer uc.osuApi.ResetStats()
	return uc.CreateTrackAndLogRecords(ctx, startTime, time.Since(startTime))
}

func (uc *UseCase) TrackSingleFollowing(ctx context.Context, following *model.Following) error {
	// check if following was not fetched in the last 24 hours (just in case)
	if err := uc.txm.ReadOnly(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		fw, _ := uc.following.Get(ctx, tx, following.ID)
		if fw == nil {
			return fmt.Errorf("following not found in db")
		}

		// actually check if it was fetched in last 23 hours cause of tracker timing error
		if time.Since(fw.LastFetched) < 23*time.Hour {
			return fmt.Errorf("user was fetched within the last 23 hours, no need to track")
		}

		return nil
	}); err != nil {
		return err
	}

	// get data from api
	user, userMapsets, err := uc.osuApi.GetUserWithMapsets(ctx, strconv.Itoa(following.ID))
	if err != nil {
		return fmt.Errorf("failed to get info from api, user id: %v, err: %w", following.ID, err)
	}

	if err = uc.createOrUpdateData(ctx, following, user, userMapsets); err != nil {
		return fmt.Errorf("failed to create or update data, user id: %v, err: %w", following.ID, err)
	}
	return nil
}
