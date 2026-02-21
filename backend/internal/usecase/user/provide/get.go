package userprovide

import (
	"context"
	"osu-dashboard/internal/database/model"
	"osu-dashboard/internal/database/txmanager"
	"osu-dashboard/internal/dto"
	"osu-dashboard/internal/usecase/mappers"
	"strconv"
)

func (uc *UseCase) Get(ctx context.Context, id int) (*dto.User, error) {
	var user *model.User
	var userExists bool

	// figure out if we have requested user in database if no then fetch osuapi
	txErr := uc.txm.ReadOnly(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		var err error
		userExists, err = uc.user.Exists(ctx, tx, id)
		if err != nil {
			return err
		}
		if !userExists {
			return nil
		}

		user, err = uc.user.Get(ctx, tx, id)
		if err != nil {
			return err
		}

		return nil
	})
	if txErr != nil {
		return nil, txErr
	}

	var userDto *dto.User
	if !userExists {
		apiUser, err := uc.osuApi.GetUser(ctx, strconv.Itoa(id))
		if err != nil {
			return nil, err
		}

		userDto, err = MapOsuApiUserToUserDTO(apiUser)
		if err != nil {
			return nil, err
		}
	} else {
		var err error
		userDto, err = mappers.MapUserModelToUserDTO(user)
		if err != nil {
			return nil, err
		}
	}

	mappers.KeepLastNKeyValuesFromStats(userDto.UserStats, statsMaxElements)
	return userDto, nil
}

func (uc *UseCase) GetByName(ctx context.Context, name string) (*dto.User, error) {
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
