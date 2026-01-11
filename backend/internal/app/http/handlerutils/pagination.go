package handlerutils

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetPageQueryParam(c echo.Context) (int, error) {
	page := c.QueryParam("page")
	var pageInt int
	if page == "" {
		pageInt = 1
	} else {
		var err error
		pageInt, err = strconv.Atoi(page)
		if err != nil {
			return 0, fmt.Errorf("invalid page param %w", err)
		}
		if pageInt <= 0 {
			return 0, fmt.Errorf("invalid page %v", page)
		}
	}

	return pageInt, nil
}
