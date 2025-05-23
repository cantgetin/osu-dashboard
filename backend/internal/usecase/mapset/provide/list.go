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

type ListResponse struct {
	Mapsets     []*dto.Mapset
	CurrentPage int
	Pages       int
}

func (uc *UseCase) List(
	ctx context.Context,
	cmd *ListCommand,
) (*ListResponse, error) {
	var dtoMapsets []*dto.Mapset
	var count int

	txErr := uc.txm.ReadOnly(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		mapsets, c, err := uc.mapset.ListWithFilterSortLimitOffset(
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
			count = c
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

	return &ListResponse{
		Mapsets:     dtoMapsets,
		CurrentPage: cmd.Page,
		Pages:       (count / mapsetsPerPage) + 1,
	}, nil
}

func (uc *UseCase) ListForUser(
	ctx context.Context,
	userID int,
	cmd *ListCommand,
) (*ListResponse, error) {
	var dtoMapsets []*dto.Mapset
	var count int

	txErr := uc.txm.ReadOnly(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		mapsets, c, err := uc.mapset.ListForUserWithFilterSortLimitOffset(
			ctx,
			tx,
			userID,
			cmd.Filter,
			cmd.Sort,
			mapsetsPerPage,
			(cmd.Page-1)*mapsetsPerPage,
		)
		if err != nil {
			return err
		}
		count = c

		mapsetIDs := make([]int, len(mapsets))
		for i, m := range mapsets {
			mapsetIDs[i] = m.ID
		}

		beatmaps, err := uc.beatmap.ListForMapsets(ctx, tx, mapsetIDs...)
		if err != nil {
			return err
		}

		// now attach beatmaps to mapsets
		for _, m := range mapsets {
			mapsetBeatmaps := make([]*model.Beatmap, 0)
			for _, bm := range beatmaps {
				if bm.MapsetID == m.ID {
					mapsetBeatmaps = append(mapsetBeatmaps, bm)
				}
			}

			dtoMapset, err := mappers.MapMapsetModelToMapsetDTO(m, mapsetBeatmaps)
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

	return &ListResponse{
		Mapsets:     dtoMapsets,
		CurrentPage: cmd.Page,
		Pages:       (count / mapsetsPerPage) + 1,
	}, nil
}
