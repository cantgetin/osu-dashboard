package trackingserviceapi

import (
	"github.com/labstack/echo/v4"
	"playcount-monitor-backend/internal/database/repository/model"
)

func (s *ServiceImpl) Create(c echo.Context) error {
	tracking := new(model.Tracking)
	if err := c.Bind(tracking); err != nil {
		return err
	}

	return s.trackingCreator.Create(c.Request().Context(), tracking)
}

func (s *ServiceImpl) List(c echo.Context) error {
	trackingList, err := s.trackingProvider.List(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(200, trackingList)
}
