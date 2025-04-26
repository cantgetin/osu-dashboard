package track

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
	"playcount-monitor-backend/internal/service/osuapi"
	"strconv"
	"time"
)

func (uc *UseCase) Track(
	ctx context.Context,
	lg *log.Logger,
) error {
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

	lg.Infof("got following IDs from db, %v total", len(follows))
	if len(follows) == 0 {
		return fmt.Errorf("no following users present in db")
	}

	// max 300 requests a minute
	for i, following := range follows {
		lg.Infof("fetching user %s with id %v, %v/%v", following.Username, following.ID, i, len(follows))
		if err := uc.TrackSingleFollowing(ctx, following); err != nil {
			return fmt.Errorf("failed to fetch specific user: %w", err)
		}
	}

	elapsed := time.Since(startTime)
	reqs := uc.osuApi.GetOutgoingRequestCount()
	avgReqsPerMin := float64(reqs) / elapsed.Minutes()

	lg.Infof("Sent %v requests to api in %v minutes", reqs, elapsed.Minutes())
	lg.Infof("Average requests per minute: %f", avgReqsPerMin)

	uc.osuApi.ResetOutgoingRequestCount()

	return nil
}

func (uc *UseCase) TrackSingleFollowing(ctx context.Context, following *model.Following) error {
	var dbUserMapsets []*model.Mapset
	if err := uc.txm.ReadOnly(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		var err error
		if dbUserMapsets, err = uc.mapset.ListForUser(ctx, tx, following.ID); err != nil {
			return err
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

	// if mapset doesnt exist in db or doesnt have genre/lang data, fetch Extended info
	for _, mapset := range userMapsets {
		// if mapset does not exist in db, then get Extended info like genre, language
		if !containsMapset(dbUserMapsets, mapset.Id) {
			langGenreInfo, err := uc.osuApi.GetMapsetExtended(ctx, strconv.Itoa(mapset.Id))
			if err != nil {
				return fmt.Errorf("failed to get mapset extended info from api, mapset id: %v, err: %w", mapset.Id, err)
			}

			mapset.Genre = langGenreInfo.Genre.Name
			mapset.Language = langGenreInfo.Language.Name
		} else {
			//if mapset exist in db but doesnt have genre/lang data, then also fetch Extended info
			dbMapset := getMapsetByID(dbUserMapsets, mapset.Id)
			if dbMapset.Genre == "" || dbMapset.Language == "" {
				langGenreInfo, err := uc.osuApi.GetMapsetExtended(ctx, strconv.Itoa(mapset.Id))
				if err != nil {
					return fmt.Errorf("failed to get mapset extended info from api, mapset id: %v, err: %w", mapset.Id, err)
				}

				mapset.Genre = langGenreInfo.Genre.Name
				mapset.Language = langGenreInfo.Language.Name
			}
		}
	}

	if err := uc.createOrUpdateData(ctx, following, user, userMapsets); err != nil {
		return fmt.Errorf("failed to create or update data, user id: %v, err: %w", following.ID, err)
	}
	return nil
}

func (uc *UseCase) createOrUpdateData(
	ctx context.Context,
	following *model.Following,
	user *osuapi.User,
	userMapsets []*osuapi.MapsetExtended,
) error {
	// create/update data in db
	txErr := uc.txm.ReadWrite(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		userExists, err := uc.user.Exists(ctx, tx, user.ID)
		if err != nil {
			return err
		}

		if userExists {
			// update
			err := uc.updateUserCard(ctx, tx, user, userMapsets)
			if err != nil {
				return fmt.Errorf("failed to update user card, user id: %v, err: %w", user.ID, err)
			}
		} else {
			// create
			err := uc.createUserCard(ctx, tx, user, userMapsets)
			if err != nil {
				return fmt.Errorf("failed to create user card, user id: %v, err: %w", user.ID, err)
			}
		}

		err = uc.following.SetLastFetchedForUser(ctx, tx, following.Username, time.Now().UTC())
		if err != nil {
			return fmt.Errorf("failed to set last fetched for following %v: %w", following.Username, err)
		}

		return nil
	})
	if txErr != nil {
		return txErr
	}

	return nil
}
