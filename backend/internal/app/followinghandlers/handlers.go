package followinghandlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Handlers) Create(c echo.Context) error {
	code := c.Param("code")
	if code == "" {
		return echo.ErrBadRequest
	}

	err := s.followingCreator.Create(c.Request().Context(), code)
	if err != nil {
		s.lg.Printf("failed to create following: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create following")
	}

	return c.NoContent(http.StatusCreated)
}

func (s *Handlers) List(c echo.Context) error {
	trackingList, err := s.followingProvider.List(c.Request().Context())
	if err != nil {
		s.lg.Printf("failed to list followings: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to list followings")
	}

	return c.JSON(200, trackingList)
}
