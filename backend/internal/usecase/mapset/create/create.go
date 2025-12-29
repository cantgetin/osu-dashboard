package mapsetcreate

import (
	"context"
	"fmt"
	"osu-dashboard/internal/database/repository/model"
	"osu-dashboard/internal/database/txmanager"
	"osu-dashboard/internal/usecase/command"
	"osu-dashboard/internal/usecase/mappers"
)

func (uc *UseCase) Create(ctx context.Context, cmd *command.CreateMapsetCommand) error {
	txErr := uc.txm.ReadWrite(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		mapsetExists, err := uc.mapset.Exists(ctx, tx, cmd.Id)
		if err != nil {
			return err
		}

		if mapsetExists {
			return fmt.Errorf("mapset with id %v already exists", cmd.Id)
		}

		// create mapset
		mapset, err := mappers.MapCreateMapsetCommandToMapsetModel(cmd)
		if err != nil {
			return err
		}

		err = uc.mapset.Create(ctx, tx, mapset)
		if err != nil {
			return err
		}

		// create beatmaps
		for _, beatmap := range cmd.Beatmaps {
			var beatmapModel *model.Beatmap
			beatmapModel, err = mappers.MapCreateBeatmapCommandToBeatmapModel(beatmap)
			if err != nil {
				return err
			}

			err = uc.beatmap.Create(ctx, tx, beatmapModel)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if txErr != nil {
		return txErr
	}

	return nil
}
