package usercreate

import (
	"context"
	"fmt"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
	"time"
)

func (uc *UseCase) Create(
	ctx context.Context,
	cmd *CreateUserCommand,
) error {
	txErr := uc.txm.ReadWrite(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		userExists, err := uc.user.Exists(ctx, tx, cmd.ID)
		if err != nil {
			return err
		}

		if userExists {
			return fmt.Errorf("user with id %v already exists", cmd.ID)
		}

		// create user
		user, err := mapCommandToUserModel(cmd)
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

func mapCommandToUserModel(cmd *CreateUserCommand) (*model.User, error) {
	return &model.User{
		ID:                       cmd.ID,
		Username:                 cmd.Username,
		AvatarURL:                cmd.AvatarURL,
		GraveyardBeatmapsetCount: cmd.GraveyardBeatmapsetCount,
		UnrankedBeatmapsetCount:  cmd.UnrankedBeatmapsetCount,
		CreatedAt:                time.Now().UTC(),
		UpdatedAt:                time.Now().UTC(),
	}, nil
}
