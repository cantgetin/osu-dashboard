package mapsetserviceapi

import (
	"github.com/labstack/echo/v4"
	"playcount-monitor-backend/internal/database/repository/model"
)

func (s *ServiceImpl) Create(c echo.Context) error {
	mapset := new(model.Mapset)
	if err := c.Bind(mapset); err != nil {
		return err
	}

	return s.mapsetProvider.Create(c.Request().Context(), mapset)
}
