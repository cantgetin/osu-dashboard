package statistichandlers

import (
	"net/http"

	"osu-dashboard/internal/app/http/handlerutils"

	"github.com/labstack/echo/v4"
)

func (h *Handlers) GetUserMapStatistics(c echo.Context) error {
	id, err := handlerutils.GetIdQueryParam(c)
	if err != nil {
		return err
	}

	userStatistics, err := h.statisticProvider.GetForUser(c.Request().Context(), id)
	if err != nil {
		return handlerutils.EchoInternalError(err)
	}

	return c.JSON(http.StatusOK, userStatistics)
}

func (h *Handlers) GetSystemStatistics(c echo.Context) error {
	systemStats, err := h.statisticProvider.GetForSystem(c.Request().Context())
	if err != nil {
		return handlerutils.EchoInternalError(err)
	}

	return c.JSON(http.StatusOK, systemStats)
}
