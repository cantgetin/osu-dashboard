package pingserviceapi

import (
	"github.com/labstack/echo/v4"
)

func (s *ServiceImpl) Ping(c echo.Context) error {
	return c.JSON(200, "pong")
}
