package mapsethandlers

import (
	"net/http"
	handlerutils2 "osu-dashboard/internal/app/http/handlerutils"
	"osu-dashboard/internal/usecase/command"
	mapsetprovide "osu-dashboard/internal/usecase/mapset/provide"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handlers) Create(c echo.Context) error {
	mapset := new(command.CreateMapsetCommand)
	if err := c.Bind(mapset); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.mapsetCreator.Create(c.Request().Context(), mapset); err != nil {
		return handlerutils2.EchoInternalError(err)
	}

	return c.JSON(http.StatusCreated, mapset)
}

func (h *Handlers) Get(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "empty id param")
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	mapset, err := h.mapsetProvider.Get(c.Request().Context(), idInt)
	if err != nil {
		return handlerutils2.EchoInternalError(err)
	}

	return c.JSON(http.StatusOK, mapset)
}

func (h *Handlers) List(c echo.Context) error {
	pageInt, err := handlerutils2.GetPageQueryParam(c)
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

	listResp, err := h.mapsetProvider.List(
		c.Request().Context(),
		&mapsetprovide.ListCommand{
			Page:   pageInt,
			Sort:   mapsetSort,
			Filter: mapsetFilter,
		},
	)
	if err != nil {
		return handlerutils2.EchoInternalError(err)
	}

	return c.JSON(http.StatusOK, listResp)
}

func (h *Handlers) ListForUser(c echo.Context) error {
	idInt, err := getUserIDFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	pageInt, err := handlerutils2.GetPageQueryParam(c)
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

	listResp, err := h.mapsetProvider.ListForUser(
		c.Request().Context(),
		idInt,
		&mapsetprovide.ListCommand{
			Page:   pageInt,
			Sort:   mapsetSort,
			Filter: mapsetFilter,
		},
	)
	if err != nil {
		return handlerutils2.EchoInternalError(err)
	}

	return c.JSON(http.StatusOK, listResp)
}
