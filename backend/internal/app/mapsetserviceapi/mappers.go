package mapsetserviceapi

import (
	"errors"
	"github.com/labstack/echo/v4"
	"playcount-monitor-backend/internal/database/repository/model"
	"strconv"
)

func mapSortQueryParamsToMapsetSort(fieldParam string, directionParam string) model.MapsetSort {
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

func mapSearchAndFilterQueryParamsToMapsetFilter(search string, status string) model.MapsetFilter {
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

func getPageQueryParam(c echo.Context) (int, error) {
	page := c.QueryParam("page")
	var pageInt int
	if page == "" {
		pageInt = 1
	} else {
		var err error
		pageInt, err = strconv.Atoi(page)
		if err != nil || pageInt <= 0 {
			return 0, err
		}
	}

	return pageInt, nil
}

func getUserIDFromContext(c echo.Context) (int, error) {
	id := c.Param("id")
	if id == "" {
		return 0, errors.New("invalid user ID")
	}
	return strconv.Atoi(id)
}
