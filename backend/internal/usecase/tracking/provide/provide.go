package trackingprovide

import (
	"context"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
	"playcount-monitor-backend/internal/dto"
)

func (uc *UseCase) List(
	ctx context.Context,
) ([]*dto.Tracking, error) {
	var trackList []*model.Tracking
	txErr := uc.txm.ReadOnly(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		var err error
		trackList, err = uc.tracking.List(ctx, tx)
		if err != nil {
			return err
		}

		return nil
	})
	if txErr != nil {
		return nil, txErr
	}

	return mapTrackingModelToTrackingDTO(trackList), nil
}

func mapTrackingModelToTrackingDTO(trackList []*model.Tracking) []*dto.Tracking {
	var trackDTOList []*dto.Tracking
	for _, track := range trackList {
		trackDTOList = append(trackDTOList, &dto.Tracking{
			ID:            track.ID,
			Username:      track.Username,
			TrackingSince: track.CreatedAt,
		})
	}

	return trackDTOList
}
