package mapsetserviceapi

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"playcount-monitor-backend/internal/usecase/command"
	mapsetprovide "playcount-monitor-backend/internal/usecase/mapset/provide"
	"strconv"
)

func (s *ServiceImpl) Create(c echo.Context) error {
	mapset := new(command.CreateMapsetCommand)
	if err := c.Bind(mapset); err != nil {
		return err
	}

	return s.mapsetCreator.Create(c.Request().Context(), mapset)
}

func (s *ServiceImpl) Get(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return echo.ErrBadRequest
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return echo.ErrBadRequest
	}

	mapset, err := s.mapsetProvider.Get(c.Request().Context(), idInt)

	if err != nil {
		return err
	}

	return c.JSON(200, mapset)
}

func (s *ServiceImpl) List(c echo.Context) error {
	pageInt, err := getPageQueryParam(c)
	if err != nil {
		return echo.ErrBadRequest
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
		return c.JSON(http.StatusInternalServerError, err)
	}

	response := MapsetListResponse{
		Mapsets:     listResp.Mapsets,
		CurrentPage: listResp.CurrentPage,
		Pages:       listResp.Pages,
	}

	return c.JSON(http.StatusOK, response)
}

func (s *ServiceImpl) ListForUser(c echo.Context) error {
	idInt, err := getUserIDFromContext(c)
	if err != nil {
		return echo.ErrBadRequest
	}

	pageInt, err := getPageQueryParam(c)
	if err != nil {
		return echo.ErrBadRequest
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
		return c.JSON(http.StatusInternalServerError, err)
	}

	response := MapsetListResponse{
		Mapsets:     listResp.Mapsets,
		CurrentPage: listResp.CurrentPage,
		Pages:       listResp.Pages,
	}

	return c.JSON(http.StatusOK, response)
}
