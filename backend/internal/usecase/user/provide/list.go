package userprovide

import (
	"context"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
	"playcount-monitor-backend/internal/dto"
	"playcount-monitor-backend/internal/usecase/mappers"
	"sort"
)

const (
	statsMaxElements = 7
	usersPerPage     = 50
)

func (uc *UseCase) ListOld(
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

	sort.Slice(outUsers, func(i, j int) bool {
		var pc1, pc2 int
		for _, v := range outUsers[i].UserStats {
			pc1 = v.PlayCount
			break
		}
		for _, v := range outUsers[j].UserStats {
			pc2 = v.PlayCount
			break
		}
		return pc1 > pc2
	})

	return outUsers, nil
}

type ListIn struct {
	Page   int
	Sort   model.UserSort
	Filter model.UserFilter
}

type ListOut struct {
	Users       []*dto.User
	CurrentPage int
	Pages       int
}

func (uc *UseCase) List(
	ctx context.Context,
	cmd *ListIn,
) (*ListOut, error) {
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

	return &ListOut{
		Users:       outUsers,
		CurrentPage: cmd.Page,
		Pages:       (count / usersPerPage) + 1,
	}, nil
}
