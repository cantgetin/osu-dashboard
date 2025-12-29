package userhandlers

import (
	"net/http"
	"osu-dashboard/internal/app/handlerutils"
	"osu-dashboard/internal/database/repository/model"
	"osu-dashboard/internal/dto"
	userprovide "osu-dashboard/internal/usecase/user/provide"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (s *Handlers) Create(c echo.Context) error {
	user := new(dto.User)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := s.userCreator.Create(c.Request().Context(), user); err != nil {
		return handlerutils.EchoInternalError(err)
	}
	return c.JSON(http.StatusCreated, user)
}

func (s *Handlers) Update(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := s.userUpdater.Update(c.Request().Context(), user); err != nil {
		return handlerutils.EchoInternalError(err)
	}
	return c.JSON(http.StatusAccepted, user)
}

func (s *Handlers) Get(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "empty user id")
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := s.userProvider.Get(c.Request().Context(), idInt)
	if err != nil {
		return handlerutils.EchoInternalError(err)
	}

	return c.JSON(http.StatusOK, user)
}

func (s *Handlers) GetByName(c echo.Context) error {
	name := c.Param("name")
	user, err := s.userProvider.GetByName(c.Request().Context(), name)
	if err != nil {
		return handlerutils.EchoInternalError(err)
	}

	return c.JSON(http.StatusOK, user)
}

func (s *Handlers) List(c echo.Context) error {
	pageInt, err := handlerutils.GetPageQueryParam(c)
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

	users, err := s.userProvider.List(c.Request().Context(), &userprovide.ListIn{
		Page:   pageInt,
		Sort:   userSort,
		Filter: userFilter,
	})
	if err != nil {
		return handlerutils.EchoInternalError(err)
	}

	return c.JSON(http.StatusOK, users)
}
