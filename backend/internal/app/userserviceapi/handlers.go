package userserviceapi

import (
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/dto"
	userprovide "playcount-monitor-backend/internal/usecase/user/provide"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (s *ServiceImpl) Create(c echo.Context) error {
	user := new(dto.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	return s.userCreator.Create(c.Request().Context(), user)
}

func (s *ServiceImpl) Update(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	return s.userUpdater.Update(c.Request().Context(), user)
}

func (s *ServiceImpl) Get(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return echo.ErrBadRequest
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return echo.ErrBadRequest
	}

	user, err := s.userProvider.Get(c.Request().Context(), idInt)
	if err != nil {
		return err
	}

	return c.JSON(200, user)
}

func (s *ServiceImpl) GetByName(c echo.Context) error {
	name := c.Param("name")
	user, err := s.userProvider.GetByName(c.Request().Context(), name)
	if err != nil {
		return err
	}

	return c.JSON(200, user)
}

func (s *ServiceImpl) List(c echo.Context) error {
	pageInt, err := getPageQueryParam(c)
	if err != nil {
		return echo.ErrBadRequest
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
		return err
	}

	return c.JSON(200, users)
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
