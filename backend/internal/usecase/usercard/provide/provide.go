package usercardprovide

import (
	"context"
	"playcount-monitor-backend/internal/database/txmanager"
	"playcount-monitor-backend/internal/dto"
	"playcount-monitor-backend/internal/usecase/mappers"
)

const mapsetsPerPage = 50
const statsMaxElements = 7

func (uc *UseCase) Get(
	ctx context.Context,
	userID int,
	page int,
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

		// get user map counts
		statuses, err := uc.mapset.ListStatusesForUser(ctx, tx, userID)
		userCard.User.UserMapCounts = mappers.MapStatusesToUserMapCounts(statuses)

		// get user mapsets
		mapsets, err := uc.mapset.ListForUserWithLimitOffset(ctx, tx, userID, mapsetsPerPage, (page-1)*mapsetsPerPage)
		if err != nil {
			return err
		}

		// for each mapset get its beatmaps and map to DTO
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

	KeepLastNStatsValuesFromUserCard(userCard, statsMaxElements)
	return userCard, nil
}

func KeepLastNStatsValuesFromUserCard(userCard *dto.UserCard, n int) {
	mappers.KeepLastNKeyValuesFromStats(userCard.User.UserStats, n)

	for _, mapset := range userCard.Mapsets {
		mappers.KeepLastNKeyValuesFromStats(mapset.MapsetStats, n)
		for _, beatmap := range mapset.Beatmaps {
			mappers.KeepLastNKeyValuesFromStats(beatmap.BeatmapStats, n)
		}
	}
}
