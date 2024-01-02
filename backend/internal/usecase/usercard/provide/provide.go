package usercardprovide

import (
	"context"
	"playcount-monitor-backend/internal/database/txmanager"
	"playcount-monitor-backend/internal/dto"
	"playcount-monitor-backend/internal/usecase/mappers"
)

func (uc *UseCase) Get(
	ctx context.Context,
	userID int,
) (*dto.UserCard, error) {
	var userCard = new(dto.UserCard)
	txErr := uc.txm.ReadOnly(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		// get user
		user, err := uc.user.Get(ctx, tx, userID)
		if err != nil {
			return err
		}

		userCard.User, err = mappers.MapUserModelToUserDTO(user)
		if err != nil {
			return err
		}

		// get user mapsets
		mapsets, err := uc.mapset.ListForUser(ctx, tx, userID)
		if err != nil {
			return err
		}

		// for each mapset get its maps and map to DTO
		for _, mapset := range mapsets {
			beatmaps, err := uc.beatmap.ListForMapset(ctx, tx, mapset.ID)
			if err != nil {
				return err
			}

			mapsetWithMaps, err := mappers.MapMapsetModelToMapsetDTO(mapset, beatmaps)
			if err != nil {
				return err
			}
			userCard.Mapsets = append(userCard.Mapsets, mapsetWithMaps)
		}

		return nil
	})
	if txErr != nil {
		return nil, txErr
	}

	return userCard, nil
}
