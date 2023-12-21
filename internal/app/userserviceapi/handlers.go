package userserviceapi

import (
	"playcount-monitor-backend/internal/database/repository/model"
	usercreate "playcount-monitor-backend/internal/usecase/user/create"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (s *ServiceImpl) Create(c echo.Context) error {
	user := new(usercreate.CreateUserCommand)
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
	users, err := s.userProvider.List(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(200, users)
}
