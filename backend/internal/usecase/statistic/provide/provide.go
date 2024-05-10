package statisticprovide

import (
	"context"
	"playcount-monitor-backend/internal/database/txmanager"
	"playcount-monitor-backend/internal/dto"
	"strings"
)

func (uc *UseCase) Get(
	ctx context.Context,
	userID int,
) (*dto.Mapset, error) {
	tags := make(map[string]int)

	txErr := uc.txm.ReadOnly(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		mapsets, err := uc.mapset.ListForUser(ctx, tx, userID)
		if err != nil {
			return err
		}

		for _, mapset := range mapsets {
			tagsArr := strings.Fields(mapset.Tags)
			for _, tag := range tagsArr {
				tags[tag]++
			}
		}

		//beatmaps, err := uc.beatmap.ListForMapset(ctx, tx, id)
		//if err != nil {
		//	return err
		//}
		//
		//dtoMapset, err = mappers.MapMapsetModelToMapsetDTO(mapset, beatmaps)
		//if err != nil {
		//	return err
		//}

		//return nil
		return nil
	})
	if txErr != nil {
		return nil, txErr
	}

	//mappers.KeepLastNKeyValuesFromStats(dtoMapset.MapsetStats, statsMaxElements)
	//for _, beatmap := range dtoMapset.Beatmaps {
	//	mappers.KeepLastNKeyValuesFromStats(beatmap.BeatmapStats, statsMaxElements)
	//}

	return nil, nil
}
