package searchusecase

import (
	"context"
	"fmt"
	"osu-dashboard/internal/database/repository/model"
	"osu-dashboard/internal/database/txmanager"
	"osu-dashboard/internal/dto"
)

const searchItemLimit = 3

func (uc *UseCase) Search(ctx context.Context, query string) (result []*dto.SearchResult, err error) {
	result = make([]*dto.SearchResult, 0)

	// search users first
	var users []*model.User
	usersFilter := model.UserFilter{model.UserNameField: query}

	if txErr := uc.txm.ReadOnly(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		users, err = uc.user.ListUsersWithFilterAndLimit(ctx, tx, usersFilter, searchItemLimit)
		if err != nil {
			return fmt.Errorf("failed to list users with filter and limit: %w", err)
		}

		return nil
	}); txErr != nil {
		return nil, txErr
	}

	// search mapsets second
	var mapsets []*model.Mapset
	mapsetsFilter := model.MapsetFilter{model.MapsetArtistOrTitleOrTagsFields: query}

	if txErr := uc.txm.ReadOnly(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		mapsets, err = uc.mapset.ListWithFilterAndLimit(ctx, tx, mapsetsFilter, searchItemLimit)
		if err != nil {
			return fmt.Errorf("failed to list mapsets with filter and limit: %w", err)
		}

		return nil
	}); txErr != nil {
		return nil, txErr
	}

	// fill search results
	for _, user := range users {
		result = append(result, &dto.SearchResult{
			Title:      user.Username,
			PictureURL: user.AvatarURL,
			Type:       dto.UserResult,
		})
	}
	for _, mapset := range mapsets {
		result = append(result, &dto.SearchResult{
			Title:      fmt.Sprintf("%s - %s", mapset.Artist, mapset.Title),
			PictureURL: mapset.PreviewURL,
			Type:       dto.MapsetResult,
		})
	}
	
	return result, nil
}
