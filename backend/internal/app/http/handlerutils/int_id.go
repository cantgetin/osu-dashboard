package handlerutils

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetIdQueryParam(c echo.Context) (int, error) {
	id := c.Param("id")
	if id == "" {
		return 0, echo.NewHTTPError(http.StatusBadRequest, "empty user id")
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return 0, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return idInt, nil
}
