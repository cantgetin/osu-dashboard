package mapsethandlers

import (
	"net/http"
	"osu-dashboard/internal/app/handlerutils"
	"osu-dashboard/internal/usecase/command"
	mapsetprovide "osu-dashboard/internal/usecase/mapset/provide"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (s *Handlers) Create(c echo.Context) error {
	mapset := new(command.CreateMapsetCommand)
	if err := c.Bind(mapset); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := s.mapsetCreator.Create(c.Request().Context(), mapset); err != nil {
		return handlerutils.EchoInternalError(err)
	}

	return c.JSON(http.StatusCreated, mapset)
}

func (s *Handlers) Get(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "empty id param")
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	mapset, err := s.mapsetProvider.Get(c.Request().Context(), idInt)
	if err != nil {
		return handlerutils.EchoInternalError(err)
	}

	return c.JSON(http.StatusOK, mapset)
}

func (s *Handlers) List(c echo.Context) error {
	pageInt, err := handlerutils.GetPageQueryParam(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	mapsetSort := mapSortQueryParamsToMapsetSort(
		c.QueryParam("sort"),
		c.QueryParam("direction"),
	)

	mapsetFilter := mapSearchAndFilterQueryParamsToMapsetFilter(
		c.QueryParam("search"),
		c.QueryParam("status"),
	)

	listResp, err := s.mapsetProvider.List(
		c.Request().Context(),
		&mapsetprovide.ListCommand{
			Page:   pageInt,
			Sort:   mapsetSort,
			Filter: mapsetFilter,
		},
	)
	if err != nil {
		return handlerutils.EchoInternalError(err)
	}

	return c.JSON(http.StatusOK, listResp)
}

func (s *Handlers) ListForUser(c echo.Context) error {
	idInt, err := getUserIDFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	pageInt, err := handlerutils.GetPageQueryParam(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	mapsetSort := mapSortQueryParamsToMapsetSort(
		c.QueryParam("sort"),
		c.QueryParam("direction"),
	)

	mapsetFilter := mapSearchAndFilterQueryParamsToMapsetFilter(
		c.QueryParam("search"),
		c.QueryParam("status"),
	)

	listResp, err := s.mapsetProvider.ListForUser(
		c.Request().Context(),
		idInt,
		&mapsetprovide.ListCommand{
			Page:   pageInt,
			Sort:   mapsetSort,
			Filter: mapsetFilter,
		},
	)
	if err != nil {
		return handlerutils.EchoInternalError(err)
	}

	return c.JSON(http.StatusOK, listResp)
}
