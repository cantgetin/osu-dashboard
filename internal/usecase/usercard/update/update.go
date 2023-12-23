package usercardupdate

import (
	"context"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
	"playcount-monitor-backend/internal/usecase/mappers"
)

func (uc *UseCase) Update(
	ctx context.Context,
	cmd *UpdateUserCardCommand,
) error {
	txErr := uc.txm.ReadWrite(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		// update user
		user, err := mappers.MapUserDTOToUserModel(cmd.User)
		if err != nil {
			return err
		}

		err = uc.user.Update(ctx, tx, user)
		if err != nil {
			return err
		}

		// update user mapsets
		for _, ms := range cmd.Mapsets {
			// update mapset
			var existingMapset *model.Mapset
			existingMapset, err = uc.mapset.Get(ctx, tx, ms.Id)
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
				newBeatmap, err = mappers.MapBeatmapDTOToBeatmapModel(&bm)
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
		}

		return nil
	})
	if txErr != nil {
		return txErr
	}

	return nil
}
