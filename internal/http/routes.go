package http

func (s *Server) setupRoutes() {
	s.server.GET("/ping", s.ping.Ping)

	s.server.GET("/users/:id", s.user.Get)
	s.server.GET("/users/:name", s.user.GetByName)
	s.server.GET("/users", s.user.List)

	s.server.POST("/user_card", s.userCard.Create)
	s.server.PUT("/user_card", s.userCard.Update)
	s.server.GET("/user_card/:id", s.userCard.Get)
}
