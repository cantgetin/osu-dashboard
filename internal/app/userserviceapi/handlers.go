package userserviceapi

import (
	"github.com/labstack/echo/v4"
	"playcount-monitor-backend/internal/repository/model"
)

func (s *ServiceImpl) Create(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	return s.userProvider.Create(c.Request().Context(), user)
}

func (s *ServiceImpl) Update(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	return s.userProvider.Update(c.Request().Context(), user)
}

func (s *ServiceImpl) Get(c echo.Context) error {
	id := c.Param("id")
	user, err := s.userProvider.Get(c.Request().Context(), id)
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
