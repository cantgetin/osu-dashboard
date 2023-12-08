package beatmapserviceapi

import (
	"github.com/labstack/echo/v4"
	"playcount-monitor-backend/internal/repository/model"
)

func (s *ServiceImpl) Create(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	return s.userProvider.Create(c.Request().Context(), user)
}
