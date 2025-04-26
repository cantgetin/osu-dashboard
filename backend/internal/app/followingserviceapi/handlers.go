package followingserviceapi

import (
	"github.com/labstack/echo/v4"
)

func (s *ServiceImpl) Create(c echo.Context) error {
	code := c.Param("code")
	if code == "" {
		return echo.ErrBadRequest
	}

	return s.followingCreator.Create(c.Request().Context(), code)
}

func (s *ServiceImpl) List(c echo.Context) error {
	trackingList, err := s.followingProvider.List(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(200, trackingList)
}
