package http

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"
	"net/http"
)

func RateLimitMiddleware(r rate.Limit, b int) echo.MiddlewareFunc {
	limiter := rate.NewLimiter(r, b)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !limiter.Allow() {
				return c.JSON(http.StatusTooManyRequests, map[string]string{
					"error": "rate limit exceeded",
				})
			}
			return next(c)
		}
	}
}
