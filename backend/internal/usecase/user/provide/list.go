package userprovide

import (
	"context"
	"osu-dashboard/internal/database/repository/model"
	"osu-dashboard/internal/database/txmanager"
	"osu-dashboard/internal/dto"
	"osu-dashboard/internal/usecase/mappers"
)

const (
	statsMaxElements = 7
	usersPerPage     = 50
)

type ListIn struct {
	Page   int
	Sort   model.UserSort
	Filter model.UserFilter
}

func (uc *UseCase) List(ctx context.Context, cmd *ListIn) (*dto.UsersPaged, error) {
	var users []*model.User
	var count int

	txErr := uc.txm.ReadOnly(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		var err error
		users, count, err = uc.user.ListUsersWithFilterSortLimitOffset(
			ctx,
			tx,
			cmd.Filter,
			cmd.Sort,
			usersPerPage,
			(cmd.Page-1)*usersPerPage,
		)
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

	return &dto.UsersPaged{
		Users:       outUsers,
		CurrentPage: cmd.Page,
		Pages:       (count + usersPerPage - 1) / usersPerPage,
	}, nil
}
