package loghandlers

import (
	"net/http"
	handlerutils2 "osu-dashboard/internal/app/http/handlerutils"

	"github.com/labstack/echo/v4"
)

func (h *Handlers) List(c echo.Context) error {
	page, err := handlerutils2.GetPageQueryParam(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	logs, err := h.logProvider.List(c.Request().Context(), page)
	if err != nil {
		return handlerutils2.EchoInternalError(err)
	}

	return c.JSON(http.StatusOK, logs)
}
