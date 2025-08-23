package userserviceapi

import "playcount-monitor-backend/internal/database/repository/model"

func mapSortQueryParamsToUserSort(fieldParam string, directionParam string) model.UserSort {
	var res model.UserSort

	if fieldParam != "" && directionParam != "" {
		var field model.UserMapStatsSortFields
		var direction model.SortDirection

		switch fieldParam {
		case "playcount":
			field = model.UserPlaycount
		case "map_count":
			field = model.UserMapCount
		case "favourites":
			field = model.UserFavs
		case "comments":
			field = model.UserComms
		}

		switch directionParam {
		case "asc":
			direction = model.ASC
		case "desc":
			direction = model.DESC
		}

		res = model.UserSort{
			Field:     field,
			Direction: direction,
		}
	}

	return res
}

func mapSearchAndFilterQueryParamsToMapsetFilter(search string) model.UserFilter {
	res := make(model.UserFilter)

	if search != "" {
		res[model.UserNameField] = search
	}

	return res
}
