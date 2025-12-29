package loghandlers

import (
	"net/http"
	"osu-dashboard/internal/app/handlerutils"

	"github.com/labstack/echo/v4"
)

func (s *Handlers) List(c echo.Context) error {
	page, err := handlerutils.GetPageQueryParam(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	logs, err := s.logProvider.List(c.Request().Context(), page)
	if err != nil {
		return handlerutils.EchoInternalError(err)
	}

	return c.JSON(http.StatusOK, logs)
}
