package track

import (
	"context"
	"fmt"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
	"playcount-monitor-backend/internal/service/osuapi"
	"playcount-monitor-backend/internal/usecase/command"
	"playcount-monitor-backend/internal/usecase/mappers"
	"strconv"
)

func (uc *UseCase) Track(
	ctx context.Context,
) error {
	// get all following IDs from db and get updated data from api, update data in db

	// get IDs
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

	if len(follows) == 0 {
		return fmt.Errorf("no following users present in db")
	}

	for _, following := range follows {
		// get data from api
		user, userMapsets, err := uc.osuApiService.GetUserWithMapsets(ctx, strconv.Itoa(following.ID))
		if err != nil {
			return fmt.Errorf("failed to get info from api, user id: %v, err: %w", following.ID, err)
		}

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

			return nil
		})
		if txErr != nil {
			return txErr
		}
	}

	return nil
}

func (uc *UseCase) createUserCard(
	ctx context.Context,
	tx txmanager.Tx,
	user *osuapi.User,
	userMapsets []*osuapi.Mapset,
) error {
	cmd := &command.CreateUserCardCommand{
		User:    mapOsuApiUserToCreateUserCommand(user),
		Mapsets: mapOsuApiMapsetsToCreateMapsetCommands(userMapsets),
	}

	// create user
	userModel, err := mappers.MapCreateUserCardCommandToUserModel(cmd)
	if err != nil {
		return err
	}

	err = uc.user.Create(ctx, tx, userModel)
	if err != nil {
		return err
	}

	// create user mapsets
	for _, ms := range cmd.Mapsets {
		// create mapset
		var mapset *model.Mapset
		mapset, err = mappers.MapCreateMapsetCommandToMapsetModel(ms)
		if err != nil {
			return err
		}

		err = uc.mapset.Create(ctx, tx, mapset)
		if err != nil {
			return err
		}

		// create mapset beatmaps
		for _, bm := range ms.Beatmaps {
			var beatmap *model.Beatmap
			beatmap, err = mappers.MapCreateBeatmapCommandToBeatmapModel(bm)
			if err != nil {
				return err
			}

			err = uc.beatmap.Create(ctx, tx, beatmap)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (uc *UseCase) updateUserCard(
	ctx context.Context,
	tx txmanager.Tx,
	user *osuapi.User,
	userMapsets []*osuapi.Mapset,
) error {
	// create cmd
	cmd := &command.UpdateUserCardCommand{
		User:    mapOsuApiUserToUpdateUserCommand(user),
		Mapsets: mapOsuApiMapsetsToUpdateMapsetCommands(userMapsets),
	}

	// update user
	existingUser, err := uc.user.Get(ctx, tx, cmd.User.ID)
	if err != nil {
		return err
	}

	newUser := mappers.MapUpdateUserCardCommandToUserModel(cmd)
	newUser.CreatedAt = existingUser.CreatedAt

	// add new map entry to stats json
	newUser.UserStats, err = mappers.AppendNewUserStats(existingUser.UserStats, newUser.UserStats)
	if err != nil {
		return err
	}

	err = uc.user.Update(ctx, tx, newUser)
	if err != nil {
		return err
	}

	// update user mapsets
	for _, ms := range cmd.Mapsets {
		var mapsetExist bool
		mapsetExist, err = uc.mapset.Exists(ctx, tx, ms.Id)
		if err != nil {
			return err
		}

		if mapsetExist {
			err = uc.updateExistingMapsetAndItsBeatmaps(ctx, tx, ms)
			if err != nil {
				return err
			}
		} else {
			err = uc.createNewMapsetWithItsBeatmaps(ctx, tx, ms)
			if err != nil {
				return err
			}
		}

	}

	return nil
}

func (uc *UseCase) updateExistingMapsetAndItsBeatmaps(
	ctx context.Context,
	tx txmanager.Tx,
	ms *command.UpdateMapsetCommand,
) error {
	// update mapset
	existingMapset, err := uc.mapset.Get(ctx, tx, ms.Id)
	if err != nil {
		return err
	}

	var newMapset *model.Mapset
	newMapset, err = mappers.MapUpdateMapsetCommandToMapsetModel(ms)
	if err != nil {
		return err
	}

	newMapset.MapsetStats, err = mappers.AppendNewMapsetStats(existingMapset.MapsetStats, newMapset.MapsetStats)
	if err != nil {
		return err
	}
	newMapset.CreatedAt = existingMapset.CreatedAt

	err = uc.mapset.Update(ctx, tx, newMapset)
	if err != nil {
		return err
	}

	// update mapset beatmaps
	for _, bm := range ms.Beatmaps {
		var existingBeatmap *model.Beatmap
		existingBeatmap, err = uc.beatmap.Get(ctx, tx, bm.Id)
		if err != nil {
			return err
		}

		var newBeatmap *model.Beatmap
		newBeatmap, err = mappers.MapUpdateBeatmapCommandToBeatmapModel(bm)
		if err != nil {
			return err
		}

		newBeatmap.BeatmapStats, err = mappers.AppendNewBeatmapStats(
			existingBeatmap.BeatmapStats,
			newBeatmap.BeatmapStats,
		)
		if err != nil {
			return err
		}
		newBeatmap.CreatedAt = existingBeatmap.CreatedAt

		err = uc.beatmap.Update(ctx, tx, newBeatmap)
		if err != nil {
			return err
		}
	}

	return nil
}

func (uc *UseCase) createNewMapsetWithItsBeatmaps(
	ctx context.Context,
	tx txmanager.Tx,
	ms *command.UpdateMapsetCommand,
) error {
	var newMapset *model.Mapset
	newMapset, err := mappers.MapUpdateMapsetCommandToMapsetModel(ms)
	if err != nil {
		return err
	}

	err = uc.mapset.Create(ctx, tx, newMapset)
	if err != nil {
		return err
	}

	for _, bm := range ms.Beatmaps {
		var newBeatmap *model.Beatmap
		newBeatmap, err = mappers.MapUpdateBeatmapCommandToBeatmapModel(bm)
		if err != nil {
			return err
		}

		err = uc.beatmap.Create(ctx, tx, newBeatmap)
		if err != nil {
			return err
		}
	}

	return nil
}
