package usercardprovide

import (
	"context"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
	"playcount-monitor-backend/internal/usecase/mappers"
)

func (uc *UseCase) GetUserCard(
	ctx context.Context,
	userID int,
) (*UserCard, error) {
	var userCard *UserCard
	txErr := uc.txm.ReadOnly(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		// get user
		user, err := uc.user.Get(ctx, tx, userID)
		if err != nil {
			return err
		}

		userCard.User = mappers.MapUserModelToUserDTO(user)

		// get user mapsets
		mapsets, err := uc.mapset.ListForUser(ctx, tx, userID)
		if err != nil {
			return err
		}

		userCard.Mapsets = mappers.MapMapsetModelsToMapsetDTOs(mapsets)

		// create user
		user, err := mappers.MapUserDTOToUserModel(cmd.User)
		if err != nil {
			return err
		}

		err = uc.user.Create(ctx, tx, user)
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
				beatmap, err = mappers.MapBeatmapDTOToBeatmapModel(&bm)
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
