package mapsetprovide

import (
	"context"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
	"playcount-monitor-backend/internal/dto"
	"playcount-monitor-backend/internal/usecase/mappers"
)

const mapsetsPerPage = 50
const statsMaxElements = 7

type ListCommand struct {
	Page   int
	Sort   model.MapsetSort
	Filter model.MapsetFilter
}

func (uc *UseCase) List(
	ctx context.Context,
	cmd *ListCommand,
) ([]*dto.Mapset, error) {
	var dtoMapsets []*dto.Mapset

	txErr := uc.txm.ReadOnly(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		mapsets, err := uc.mapset.ListWithFilterSortLimitOffset(
			ctx,
			tx,
			cmd.Filter,
			cmd.Sort,
			mapsetsPerPage,
			(cmd.Page-1)*mapsetsPerPage,
		)
		if err != nil {
			return err
		}

		for _, m := range mapsets {
			bm, err := uc.beatmap.ListForMapset(ctx, tx, m.ID)
			if err != nil {
				return err
			}

			dtoMapset, err := mappers.MapMapsetModelToMapsetDTO(m, bm)
			if err != nil {
				return err
			}

			dtoMapsets = append(dtoMapsets, dtoMapset)
		}

		return nil
	})
	if txErr != nil {
		return nil, txErr
	}

	for _, mapset := range dtoMapsets {
		mappers.KeepLastNKeyValuesFromStats(mapset.MapsetStats, statsMaxElements)
		for _, beatmap := range mapset.Beatmaps {
			mappers.KeepLastNKeyValuesFromStats(beatmap.BeatmapStats, statsMaxElements)
		}
	}

	return dtoMapsets, nil
}

func (uc *UseCase) ListForUser(
	ctx context.Context,
	userID int,
) ([]*dto.Mapset, error) {
	var dtoMapsets []*dto.Mapset

	txErr := uc.txm.ReadOnly(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		mapsets, err := uc.mapset.ListForUser(ctx, tx, userID)
		if err != nil {
			return err
		}

		for _, m := range mapsets {
			bm, err := uc.beatmap.ListForMapset(ctx, tx, m.ID)
			if err != nil {
				return err
			}

			dtoMapset, err := mappers.MapMapsetModelToMapsetDTO(m, bm)
			if err != nil {
				return err
			}

			dtoMapsets = append(dtoMapsets, dtoMapset)
		}

		return nil
	})
	if txErr != nil {
		return nil, txErr
	}

	for _, mapset := range dtoMapsets {
		mappers.KeepLastNKeyValuesFromStats(mapset.MapsetStats, statsMaxElements)
		for _, beatmap := range mapset.Beatmaps {
			mappers.KeepLastNKeyValuesFromStats(beatmap.BeatmapStats, statsMaxElements)
		}
	}

	return dtoMapsets, nil
}
