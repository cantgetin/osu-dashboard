package http

func (s *Server) setupRoutes() {
	s.server.GET("/ping", s.ping.Ping)

	s.server.POST("/user_card/create", s.userCard.Create)
	s.server.POST("/user_card/update", s.userCard.Update)
	s.server.GET("/user_card/:id", s.userCard.Get)

	s.server.GET("/following/list", s.tracking.List)
	s.server.POST("/following/create", s.tracking.Create)
}
