package usercardhandlers

import (
	"net/http"
	"osu-dashboard/internal/app/handlerutils"
	"osu-dashboard/internal/usecase/command"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (s *Handlers) Create(c echo.Context) error {
	userCard := new(command.CreateUserCardCommand)
	if err := c.Bind(userCard); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := s.userCardCreator.Create(c.Request().Context(), userCard); err != nil {
		return handlerutils.EchoInternalError(err)
	}

	return c.JSON(http.StatusCreated, userCard)
}

func (s *Handlers) Get(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "empty id")
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	page, err := handlerutils.GetPageQueryParam(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userCard, err := s.userCardProvider.Get(c.Request().Context(), idInt, page)
	if err != nil {
		return handlerutils.EchoInternalError(err)
	}

	return c.JSON(http.StatusOK, userCard)
}

func (s *Handlers) Update(c echo.Context) error {
	userCard := new(command.UpdateUserCardCommand)
	if err := c.Bind(userCard); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := s.userCardUpdater.Update(c.Request().Context(), userCard); err != nil {
		return handlerutils.EchoInternalError(err)
	}

	return c.JSON(http.StatusAccepted, userCard)
}
