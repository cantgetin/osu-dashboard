package logserviceapi

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (s *ServiceImpl) List(c echo.Context) error {
	pageInt, err := getPageQueryParam(c)
	if err != nil {
		return echo.ErrBadRequest
	}

	logs, err := s.logProvider.List(c.Request().Context(), pageInt)
	if err != nil {
		s.lg.Printf("failed to list logs: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to list logs")
	}

	return c.JSON(200, logs)
}

func getPageQueryParam(c echo.Context) (int, error) {
	page := c.QueryParam("page")
	var pageInt int
	if page == "" {
		pageInt = 1
	} else {
		var err error
		pageInt, err = strconv.Atoi(page)
		if err != nil || pageInt <= 0 {
			return 0, err
		}
	}

	return pageInt, nil
}
