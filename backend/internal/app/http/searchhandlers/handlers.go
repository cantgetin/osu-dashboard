package searchhandlers

import (
	"net/http"
	"osu-dashboard/internal/app/http/handlerutils"

	"github.com/labstack/echo/v4"
)

func (h *Handlers) Search(c echo.Context) error {
	query := c.Param("query")
	if query == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "empty query param")
	}

	result, err := h.searcher.Search(c.Request().Context(), query)
	if err != nil {
		return handlerutils.EchoInternalError(err)
	}

	return c.JSON(http.StatusOK, result)
}
