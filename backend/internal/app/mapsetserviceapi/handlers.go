package mapsetserviceapi

import (
	"github.com/labstack/echo/v4"
	"playcount-monitor-backend/internal/usecase/command"
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
