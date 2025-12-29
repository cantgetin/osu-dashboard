package statistichandlers

import (
	"net/http"
	"osu-dashboard/internal/app/handlerutils"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (s *Handlers) GetUserMapStatistics(c echo.Context) error {
	userID := c.Param("id")
	if userID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "empty user id")
	}
	idInt, err := strconv.Atoi(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userStatistics, err := s.statisticProvider.GetForUser(c.Request().Context(), idInt)
	if err != nil {
		return handlerutils.EchoInternalError(err)
	}

	return c.JSON(http.StatusOK, userStatistics)
}

func (s *Handlers) GetSystemStatistics(c echo.Context) error {
	systemStats, err := s.statisticProvider.GetForSystem(c.Request().Context())
	if err != nil {
		return handlerutils.EchoInternalError(err)
	}

	return c.JSON(http.StatusOK, systemStats)
}
