package followinghandlers

import (
	"net/http"
	"osu-dashboard/internal/app/handlerutils"

	"github.com/labstack/echo/v4"
)

func (s *Handlers) Create(c echo.Context) error {
	code := c.Param("code")
	if code == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "empty code")
	}

	err := s.followingCreator.Create(c.Request().Context(), code)
	if err != nil {
		return handlerutils.EchoInternalError(err)
	}

	return c.NoContent(http.StatusCreated)
}

func (s *Handlers) List(c echo.Context) error {
	trackingList, err := s.followingProvider.List(c.Request().Context())
	if err != nil {
		return handlerutils.EchoInternalError(err)
	}

	return c.JSON(http.StatusOK, trackingList)
}
