package track

import (
	"context"
	"fmt"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
	"playcount-monitor-backend/internal/service/osuapi"
	"playcount-monitor-backend/internal/usecase/command"
	"playcount-monitor-backend/internal/usecase/mappers"
	"time"
)

func (uc *UseCase) createUserCard(
	ctx context.Context,
	tx txmanager.Tx,
	user *osuapi.User,
	userMapsets []*osuapi.MapsetExtended,
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
	newMapset.CreatedAt = time.Now().UTC()

	err = uc.mapset.Create(ctx, tx, newMapset)
	if err != nil {
		return err
	}

	for _, bm := range ms.Beatmaps {
		err = uc.createNewBeatmap(ctx, tx, bm)
		if err != nil {
			return err
		}
	}

	return nil
}

func (uc *UseCase) createNewBeatmap(
	ctx context.Context,
	tx txmanager.Tx,
	bm *command.UpdateBeatmapCommand,
) error {
	newBeatmap, err := mappers.MapUpdateBeatmapCommandToBeatmapModel(bm)
	if err != nil {
		return err
	}
	newBeatmap.CreatedAt = time.Now().UTC()

	err = uc.beatmap.Create(ctx, tx, newBeatmap)
	if err != nil {
		return fmt.Errorf("failed to create new beatmap with id %v, err: %w", newBeatmap.ID, err)
	}

	return nil
}
