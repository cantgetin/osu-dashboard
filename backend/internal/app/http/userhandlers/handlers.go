package userhandlers

import (
	"net/http"
	handlerutils2 "osu-dashboard/internal/app/http/handlerutils"
	"osu-dashboard/internal/database/repository/model"
	"osu-dashboard/internal/dto"
	userprovide "osu-dashboard/internal/usecase/user/provide"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handlers) Create(c echo.Context) error {
	user := new(dto.User)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.userCreator.Create(c.Request().Context(), user); err != nil {
		return handlerutils2.EchoInternalError(err)
	}
	return c.JSON(http.StatusCreated, user)
}

func (h *Handlers) Update(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.userUpdater.Update(c.Request().Context(), user); err != nil {
		return handlerutils2.EchoInternalError(err)
	}
	return c.JSON(http.StatusAccepted, user)
}

func (h *Handlers) Get(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "empty user id")
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := h.userProvider.Get(c.Request().Context(), idInt)
	if err != nil {
		return handlerutils2.EchoInternalError(err)
	}

	return c.JSON(http.StatusOK, user)
}

func (h *Handlers) GetByName(c echo.Context) error {
	name := c.Param("name")
	user, err := h.userProvider.GetByName(c.Request().Context(), name)
	if err != nil {
		return handlerutils2.EchoInternalError(err)
	}

	return c.JSON(http.StatusOK, user)
}

func (h *Handlers) List(c echo.Context) error {
	pageInt, err := handlerutils2.GetPageQueryParam(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userSort := mapSortQueryParamsToUserSort(
		c.QueryParam("sort"),
		c.QueryParam("direction"),
	)

	userFilter := mapSearchAndFilterQueryParamsToMapsetFilter(
		c.QueryParam("search"),
	)

	users, err := h.userProvider.List(c.Request().Context(), &userprovide.ListIn{
		Page:   pageInt,
		Sort:   userSort,
		Filter: userFilter,
	})
	if err != nil {
		return handlerutils2.EchoInternalError(err)
	}

	return c.JSON(http.StatusOK, users)
}
