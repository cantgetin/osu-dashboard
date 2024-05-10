package followingprovide

import (
	"context"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
	"playcount-monitor-backend/internal/dto"
)

func (uc *UseCase) List(
	ctx context.Context,
) ([]*dto.Following, error) {
	var trackList []*model.Following
	txErr := uc.txm.ReadOnly(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		var err error
		trackList, err = uc.following.List(ctx, tx)
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

func mapTrackingModelToTrackingDTO(followingList []*model.Following) []*dto.Following {
	var trackDTOList []*dto.Following
	for _, follow := range followingList {
		trackDTOList = append(trackDTOList, &dto.Following{
			ID:             follow.ID,
			Username:       follow.Username,
			FollowingSince: follow.CreatedAt,
			LastFetched:    follow.LastFetched,
		})
	}

	return trackDTOList
}
