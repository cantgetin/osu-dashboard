package usercreate

import (
	"context"
	"fmt"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
	"playcount-monitor-backend/internal/dto"
	"time"
)

func (uc *UseCase) Create(
	ctx context.Context,
	dto *dto.User,
) error {
	txErr := uc.txm.ReadWrite(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		userExists, err := uc.user.Exists(ctx, tx, dto.ID)
		if err != nil {
			return err
		}

		if userExists {
			return fmt.Errorf("user with id %v already exists", dto.ID)
		}

		// create user
		user, err := mapUserDTOToUserModel(dto)
		if err != nil {
			return err
		}

		err = uc.user.Create(ctx, tx, user)
		if err != nil {
			return err
		}

		return nil
	})
	if txErr != nil {
		return txErr
	}

	return nil
}

func mapUserDTOToUserModel(dto *dto.User) (*model.User, error) {
	return &model.User{
		ID:                       dto.ID,
		Username:                 dto.Username,
		AvatarURL:                dto.AvatarURL,
		GraveyardBeatmapsetCount: dto.GraveyardBeatmapsetCount,
		UnrankedBeatmapsetCount:  dto.UnrankedBeatmapsetCount,
		CreatedAt:                time.Now().UTC(),
		UpdatedAt:                time.Now().UTC(),
	}, nil
}
