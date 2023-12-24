package usercardcreate

import (
	"context"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
	"playcount-monitor-backend/internal/usecase/mappers"
)

func (uc *UseCase) Create(
	ctx context.Context,
	cmd *CreateUserCardCommand,
) error {
	txErr := uc.txm.ReadWrite(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		// create user
		user := mappers.MapCreateUserCommandToUserModel(cmd.User)

		err := uc.user.Create(ctx, tx, user)
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
			}
		}

		return nil
	})
	if txErr != nil {
		return txErr
	}

	return nil
}
