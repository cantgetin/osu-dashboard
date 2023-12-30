package usercardupdate

import (
	"context"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
	"playcount-monitor-backend/internal/usecase/mappers"
	usercardcreate "playcount-monitor-backend/internal/usecase/models"
)

func (uc *UseCase) Update(
	ctx context.Context,
	cmd *usercardcreate.UpdateUserCardCommand,
) error {
	txErr := uc.txm.ReadWrite(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		// update user
		user := mappers.MapCreateUserCommandToUserModel(cmd.User)

		err := uc.user.Update(ctx, tx, user)
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
	})
	if txErr != nil {
		return txErr
	}

	return nil
}

func (uc *UseCase) updateExistingMapsetAndItsBeatmaps(ctx context.Context, tx txmanager.Tx, ms *usercardcreate.CreateMapsetCommand) error {
	// update mapset
	existingMapset, err := uc.mapset.Get(ctx, tx, ms.Id)
	if err != nil {
		return err
	}

	var newMapset *model.Mapset
	newMapset, err = mappers.MapCreateMapsetCommandToMapsetModel(ms)
	if err != nil {
		return err
	}

	newMapset.MapsetStats, err = mappers.AppendNewMapsetStats(
		existingMapset.MapsetStats,
		newMapset.MapsetStats,
	)
	if err != nil {
		return err
	}

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
		newBeatmap, err = mappers.MapCreateBeatmapCommandToBeatmapModel(bm)
		if err != nil {
			return err
		}

		newBeatmap.BeatmapStats, err = mappers.AppendNewBeatmapStats(
			existingBeatmap.BeatmapStats,
			newBeatmap.BeatmapStats,
		)

		err = uc.beatmap.Update(ctx, tx, newBeatmap)
		if err != nil {
			return err
		}
	}

	return nil
}

func (uc *UseCase) createNewMapsetWithItsBeatmaps(ctx context.Context, tx txmanager.Tx, ms *usercardcreate.CreateMapsetCommand) error {
	var newMapset *model.Mapset
	newMapset, err := mappers.MapCreateMapsetCommandToMapsetModel(ms)
	if err != nil {
		return err
	}

	err = uc.mapset.Create(ctx, tx, newMapset)
	if err != nil {
		return err
	}

	for _, bm := range ms.Beatmaps {
		var newBeatmap *model.Beatmap
		newBeatmap, err = mappers.MapCreateBeatmapCommandToBeatmapModel(bm)
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
