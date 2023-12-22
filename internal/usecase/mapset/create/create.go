package mapsetcreate

import (
	"context"
	"fmt"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
	"playcount-monitor-backend/internal/dto"
	"playcount-monitor-backend/internal/usecase/mappers"
)

func (uc *UseCase) Create(
	ctx context.Context,
	dto *dto.Mapset,
) error {
	txErr := uc.txm.ReadWrite(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		mapsetExists, err := uc.mapset.Exists(ctx, tx, dto.Id)
		if err != nil {
			return err
		}

		if mapsetExists {
			return fmt.Errorf("mapset with id %v already exists", dto.Id)
		}

		// create mapset
		mapset, err := mappers.MapMapsetDTOToMapsetModel(dto)
		if err != nil {
			return err
		}

		err = uc.mapset.Create(ctx, tx, mapset)
		if err != nil {
			return err
		}

		// create beatmaps
		for _, beatmap := range dto.Beatmaps {
			var beatmapModel *model.Beatmap
			beatmapModel, err = mappers.MapBeatmapDTOToBeatmapModel(&beatmap)
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
