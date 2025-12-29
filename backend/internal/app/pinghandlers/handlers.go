package pinghandlers

import (
	"github.com/labstack/echo/v4"
)

func (s *Handlers) Ping(c echo.Context) error {
	return c.JSON(200, "pong")
}
