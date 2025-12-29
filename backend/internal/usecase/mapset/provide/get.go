package mapsetprovide

import (
	"context"
	"osu-dashboard/internal/database/txmanager"
	"osu-dashboard/internal/dto"
	"osu-dashboard/internal/usecase/mappers"
)

func (uc *UseCase) Get(ctx context.Context, id int) (*dto.Mapset, error) {
	var dtoMapset *dto.Mapset
	txErr := uc.txm.ReadOnly(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		mapset, err := uc.mapset.Get(ctx, tx, id)
		if err != nil {
			return err
		}

		beatmaps, err := uc.beatmap.ListForMapset(ctx, tx, id)
		if err != nil {
			return err
		}

		dtoMapset, err = mappers.MapMapsetModelToMapsetDTO(mapset, beatmaps)
		if err != nil {
			return err
		}

		return nil
	})
	if txErr != nil {
		return nil, txErr
	}

	mappers.KeepLastNKeyValuesFromStats(dtoMapset.MapsetStats, statsMaxElements)
	for _, beatmap := range dtoMapset.Beatmaps {
		mappers.KeepLastNKeyValuesFromStats(beatmap.BeatmapStats, statsMaxElements)
	}

	return dtoMapset, nil
}
