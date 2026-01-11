package usercardhandlers

import (
	"net/http"
	handlerutils2 "osu-dashboard/internal/app/http/handlerutils"
	"osu-dashboard/internal/usecase/command"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handlers) Create(c echo.Context) error {
	userCard := new(command.CreateUserCardCommand)
	if err := c.Bind(userCard); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.userCardCreator.Create(c.Request().Context(), userCard); err != nil {
		return handlerutils2.EchoInternalError(err)
	}

	return c.JSON(http.StatusCreated, userCard)
}

func (h *Handlers) Get(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "empty id")
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	page, err := handlerutils2.GetPageQueryParam(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userCard, err := h.userCardProvider.Get(c.Request().Context(), idInt, page)
	if err != nil {
		return handlerutils2.EchoInternalError(err)
	}

	return c.JSON(http.StatusOK, userCard)
}

func (h *Handlers) Update(c echo.Context) error {
	userCard := new(command.UpdateUserCardCommand)
	if err := c.Bind(userCard); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.userCardUpdater.Update(c.Request().Context(), userCard); err != nil {
		return handlerutils2.EchoInternalError(err)
	}

	return c.JSON(http.StatusAccepted, userCard)
}
