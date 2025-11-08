package usercardserviseapi

import (
	"github.com/labstack/echo/v4"
	"osu-dashboard/internal/usecase/command"
	"strconv"
)

func (s *ServiceImpl) Create(c echo.Context) error {
	userCard := new(command.CreateUserCardCommand)
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
	if err != nil {
		return echo.ErrBadRequest
	}

	page := c.QueryParam("page")
	pageInt := 1
	if page != "" {
		pageInt, err = strconv.Atoi(page)
		if err != nil {
			return echo.ErrBadRequest
		}
		if pageInt <= 0 {
			return echo.ErrBadRequest
		}
	}

	userCard, err := s.userCardProvider.Get(c.Request().Context(), idInt, pageInt)
	if err != nil {
		return err
	}

	return c.JSON(200, userCard)
}

func (s *ServiceImpl) Update(c echo.Context) error {
	userCard := new(command.UpdateUserCardCommand)
	if err := c.Bind(userCard); err != nil {
		return err
	}

	return s.userCardUpdater.Update(c.Request().Context(), userCard)
}
