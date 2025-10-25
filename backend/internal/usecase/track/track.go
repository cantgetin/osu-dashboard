package track

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
	"strconv"
	"time"
)

func (uc *UseCase) TrackAllFollowings(
	ctx context.Context,
	startTime time.Time,
	timeSinceLast time.Duration,
) error {
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
	return uc.CreateTrackAndLogRecords(ctx, startTime, timeSinceLast)
}

func (uc *UseCase) TrackSingleFollowing(ctx context.Context, following *model.Following) error {
	// check if following was not fetched in the last 24 hours (just in case)
	if err := uc.txm.ReadOnly(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		fw, _ := uc.following.Get(ctx, tx, following.ID)
		if fw == nil {
			return fmt.Errorf("following not found in db")
		}

		// if it was fetched in last 24 hours then return error
		if time.Since(fw.LastFetched) < 24*time.Hour {
			return fmt.Errorf("user was fetched within the last 24 hours, no need to track")
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

func Enrich() {
	// TODO: move to separate tracker with 1 time/2 week interval

	//var dbUserMapsets []*model.Mapset
	//if err := uc.txm.ReadOnly(ctx, func(ctx context.Context, tx txmanager.Tx) error {
	//	var err error
	//	if dbUserMapsets, err = uc.mapset.ListForUser(ctx, tx, following.ID); err != nil {
	//		return err
	//	}
	//	return nil
	//}); err != nil {
	//	return err
	//}
	//// if mapset doesnt exist in db or doesnt have genre/lang data, fetch Extended info
	//for _, mapset := range userMapsets {
	//	// if mapset does not exist in db, then get Extended info like genre, language
	//	if !containsMapset(dbUserMapsets, mapset.Id) {
	//		langGenreInfo, err := uc.osuApi.GetMapsetExtended(ctx, strconv.Itoa(mapset.Id))
	//		if err != nil {
	//			return fmt.Errorf("failed to get mapset extended info from api, mapset id: %v, err: %w", mapset.Id, err)
	//		}
	//
	//		mapset.Genre = langGenreInfo.Genre.Name
	//		mapset.Language = langGenreInfo.Language.Name
	//	} else {
	//		//if mapset exist in db but doesnt have genre/lang data, then also fetch Extended info
	//		dbMapset := getMapsetByID(dbUserMapsets, mapset.Id)
	//		if dbMapset.Genre == "" || dbMapset.Language == "" {
	//			langGenreInfo, err := uc.osuApi.GetMapsetExtended(ctx, strconv.Itoa(mapset.Id))
	//			if err != nil {
	//				return fmt.Errorf("failed to get mapset extended info from api, mapset id: %v, err: %w", mapset.Id, err)
	//			}
	//
	//			mapset.Genre = langGenreInfo.Genre.Name
	//			mapset.Language = langGenreInfo.Language.Name
	//		}
	//	}
	//}
}
