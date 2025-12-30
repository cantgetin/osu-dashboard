package mapsethandlers

import (
	"errors"
	"osu-dashboard/internal/database/repository/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

func mapSortQueryParamsToMapsetSort(fieldParam, directionParam string) model.MapsetSort {
	var res model.MapsetSort

	if fieldParam != "" && directionParam != "" {
		var field model.MapsetSortField
		var direction model.SortDirection

		switch fieldParam {
		case "last_playcount":
			field = model.MapsetPlaycount
		case "created_at":
			field = model.MapsetCreatedAt
		case "last_favorites":
			field = model.MapsetFavs
		case "last_comments":
			field = model.MapsetComms
		}

		switch directionParam {
		case "asc":
			direction = model.ASC
		case "desc":
			direction = model.DESC
		}

		res = model.MapsetSort{
			Field:     field,
			Direction: direction,
		}
	}

	return res
}

func mapSearchAndFilterQueryParamsToMapsetFilter(search, status string) model.MapsetFilter {
	res := make(model.MapsetFilter)

	if search != "" {
		res[model.MapsetArtistOrTitleOrTagsFields] = search
	}

	if checkIfStatusIsValid(status) {
		res[model.MapsetStatusField] = status
	}

	return res
}

func checkIfStatusIsValid(status string) bool {
	return status == "graveyard" ||
		status == "wip" ||
		status == "pending" ||
		status == "ranked" ||
		status == "approved" ||
		status == "qualified" ||
		status == "loved"
}

func getUserIDFromContext(c echo.Context) (int, error) {
	id := c.Param("id")
	if id == "" {
		return 0, errors.New("invalid user ID")
	}
	return strconv.Atoi(id)
}
