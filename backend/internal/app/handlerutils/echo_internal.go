package handlerutils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func EchoInternalError(err error) *echo.HTTPError {
	if err == nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
}
