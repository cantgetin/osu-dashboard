package followingserviceapi

import (
	"github.com/labstack/echo/v4"
	"strconv"
)

func (s *ServiceImpl) Create(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.ErrBadRequest
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return echo.ErrBadRequest
	}

	username := c.Param("username")
	if username == "" {
		return echo.ErrBadRequest
	}

	return s.followingCreator.Create(c.Request().Context(), idInt, username)
}

func (s *ServiceImpl) List(c echo.Context) error {
	trackingList, err := s.followingProvider.List(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(200, trackingList)
}
