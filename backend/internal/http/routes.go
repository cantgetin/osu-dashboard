package http

func (s *Server) setupRoutes() {
	s.server.GET("/ping", s.ping.Ping)

	s.server.POST("/user_card/create", s.userCard.Create)
	s.server.POST("/user_card/update", s.userCard.Update)
	s.server.GET("/user_card/:id", s.userCard.Get)

	s.server.GET("/tracking/list", s.tracking.List)
	s.server.POST("/tracking/create", s.tracking.Create)
}
