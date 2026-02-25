package followinghandlers

import (
	"net/http"
	"osu-dashboard/internal/app/http/handlerutils"
	"osu-dashboard/internal/dto"

	"github.com/labstack/echo/v4"
)

func (h *Handlers) Create(c echo.Context) error {
	code := c.Param("code")
	if code == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "empty code")
	}

	var user *dto.User
	user, err := h.followingCreator.Create(c.Request().Context(), code)
	if err != nil {
		return handlerutils.EchoInternalError(err)
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *Handlers) List(c echo.Context) error {
	trackingList, err := h.followingProvider.List(c.Request().Context())
	if err != nil {
		return handlerutils.EchoInternalError(err)
	}

	return c.JSON(http.StatusOK, trackingList)
}
