package statisticserviceapi

import (
	"github.com/labstack/echo/v4"
	"strconv"
)

func (s *ServiceImpl) GetUserMapStatistics(c echo.Context) error {
	userID := c.Param("id")
	if userID == "" {
		return echo.ErrBadRequest
	}
	idInt, err := strconv.Atoi(userID)
	if err != nil {
		return echo.ErrBadRequest
	}

	userStatistics, err := s.statisticProvider.GetForUser(c.Request().Context(), idInt)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(200, userStatistics)
}
