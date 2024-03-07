package userprovide

import (
	"context"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
	"playcount-monitor-backend/internal/dto"
	"playcount-monitor-backend/internal/usecase/mappers"
)

const statsMaxElements = 7

func (uc *UseCase) Get(
	ctx context.Context,
	id int,
) (*dto.User, error) {
	var user *model.User
	txErr := uc.txm.ReadOnly(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		var err error
		user, err = uc.user.Get(ctx, tx, id)
		if err != nil {
			return err
		}

		return nil
	})
	if txErr != nil {
		return nil, txErr
	}

	userDto, err := mappers.MapUserModelToUserDTO(user)
	if err != nil {
		return nil, err
	}

	mappers.KeepLastNKeyValuesFromStats(userDto.UserStats, statsMaxElements)
	return userDto, nil
}

func (uc *UseCase) GetByName(
	ctx context.Context,
	name string,
) (*dto.User, error) {
	var user *model.User
	txErr := uc.txm.ReadOnly(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		var err error
		user, err = uc.user.GetByName(ctx, tx, name)
		if err != nil {
			return err
		}

		return nil
	})
	if txErr != nil {
		return nil, txErr
	}

	userDto, err := mappers.MapUserModelToUserDTO(user)
	if err != nil {
		return nil, err
	}

	mappers.KeepLastNKeyValuesFromStats(userDto.UserStats, statsMaxElements)
	return userDto, nil
}

func (uc *UseCase) List(
	ctx context.Context,
) ([]*dto.User, error) {
	var users []*model.User
	txErr := uc.txm.ReadOnly(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		var err error
		users, err = uc.user.List(ctx, tx)
		if err != nil {
			return err
		}

		return nil
	})
	if txErr != nil {
		return nil, txErr
	}

	outUsers, err := mappers.MapUserModelsToUserDTOs(users)
	if err != nil {
		return nil, err
	}

	for _, user := range outUsers {
		mappers.KeepLastNKeyValuesFromStats(user.UserStats, statsMaxElements)
	}

	return outUsers, nil
}
