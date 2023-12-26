package usercardserviseapi

import (
	"github.com/labstack/echo/v4"
	usercardcreate "playcount-monitor-backend/internal/usecase/usercard/create"
	usercardupdate "playcount-monitor-backend/internal/usecase/usercard/update"
	"strconv"
)

func (s *ServiceImpl) Create(c echo.Context) error {
	userCard := new(usercardcreate.CreateUserCardCommand)
	if err := c.Bind(userCard); err != nil {
		return err
	}

	return s.userCardCreator.Create(c.Request().Context(), userCard)
}

func (s *ServiceImpl) Get(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.ErrBadRequest
	}
	idInt, err := strconv.Atoi(id)

	userCard, err := s.userCardProvider.Get(c.Request().Context(), idInt)
	if err != nil {
		return err
	}

	return c.JSON(200, userCard)
}

func (s *ServiceImpl) Update(c echo.Context) error {
	userCard := new(usercardupdate.UpdateUserCardCommand)
	if err := c.Bind(userCard); err != nil {
		return err
	}

	return s.userCardUpdater.Update(c.Request().Context(), userCard)
}
