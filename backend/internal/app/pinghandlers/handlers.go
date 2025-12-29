package pinghandlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Handlers) Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, "pong")
}
