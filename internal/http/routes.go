package http

import "github.com/labstack/echo/v4"

func (s *Server) setupRoutes() {
	s.server.GET("/users/:id", s.user.Get)
	s.server.GET("/users/:name", s.user.GetByName)
	s.server.GET("/users", s.user.List)
	s.server.POST("/users", s.user.Create)
	s.server.PUT("/users/:id", s.user.Update)

	s.server.GET("/ping", Ping)
}

func Ping(c echo.Context) error {
	return c.String(200, "pong")
}
