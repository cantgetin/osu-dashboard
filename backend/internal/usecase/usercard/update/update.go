package usercardupdate

import (
	"context"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
	"playcount-monitor-backend/internal/usecase/command"
	"playcount-monitor-backend/internal/usecase/mappers"
)

func (uc *UseCase) Update(
	ctx context.Context,
	cmd *command.UpdateUserCardCommand,
) error {
	txErr := uc.txm.ReadWrite(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		// update user
		existingUser, err := uc.user.Get(ctx, tx, cmd.User.ID)
		if err != nil {
			return err
		}

		newUser := mappers.MapUpdateUserCommandToUserModel(cmd.User)
		newUser.CreatedAt = existingUser.CreatedAt

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
	})
	if txErr != nil {
		return txErr
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
