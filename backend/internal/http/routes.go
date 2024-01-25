package http

func (s *Server) setupRoutes() {
	s.server.GET("api/ping", s.ping.Ping)

	s.server.POST("api/user_card/create", s.userCard.Create)
	s.server.POST("api/user_card/update", s.userCard.Update)
	s.server.GET("api/user_card/:id", s.userCard.Get)

	s.server.GET("api/following/list", s.tracking.List)
	s.server.POST("api/following/create", s.tracking.Create)

	s.server.GET("api/users/list", s.user.List)
}
