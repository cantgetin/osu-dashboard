package pinghandlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handlers) Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, "pong")
}
