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
	page := c.QueryParam("page")
	pageInt := 1
	if page != "" {
		var err error
		pageInt, err = strconv.Atoi(page)
		if err != nil {
			return echo.ErrBadRequest
		}
		if pageInt <= 0 {
			return echo.ErrBadRequest
		}
	}

	mapsetSort := mapSortQueryParamsToMapsetSort(
		c.QueryParam("sort"),
		c.QueryParam("direction"),
	)

	mapsetFilter := mapSearchAndFilterQueryParamsToMapsetFilter(
		c.QueryParam("search"),
		c.QueryParam("status"),
	)

	mapsetList, err := s.mapsetProvider.List(
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

	return c.JSON(http.StatusOK, mapsetList)
}

func (s *ServiceImpl) ListForUser(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return echo.ErrBadRequest
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return echo.ErrBadRequest
	}

	page := c.QueryParam("page")
	pageInt := 1
	if page != "" {
		var err error
		pageInt, err = strconv.Atoi(page)
		if err != nil {
			return echo.ErrBadRequest
		}
		if pageInt <= 0 {
			return echo.ErrBadRequest
		}
	}

	mapsetSort := mapSortQueryParamsToMapsetSort(
		c.QueryParam("sort"),
		c.QueryParam("direction"),
	)

	mapsetFilter := mapSearchAndFilterQueryParamsToMapsetFilter(
		c.QueryParam("search"),
		c.QueryParam("status"),
	)

	mapsetList, err := s.mapsetProvider.ListForUser(
		c.Request().Context(),
		idInt,
		&mapsetprovide.ListCommand{
			Page:   pageInt,
			Sort:   mapsetSort,
			Filter: mapsetFilter,
		},
	)

	if err != nil {
		return err
	}

	return c.JSON(200, mapsetList)
}
