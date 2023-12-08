package http

func (s *Server) setupRoutes() {
	s.server.GET("/users/:id", s.user.Get)
	s.server.GET("/users/:name", s.user.GetByName)
	s.server.GET("/users", s.user.List)
	s.server.POST("/users", s.user.Create)
	s.server.PUT("/users/:id", s.user.Update)
}
