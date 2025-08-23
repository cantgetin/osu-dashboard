package http

func (s *Server) setupRoutes() {
	s.server.GET("api/ping", s.ping.Ping)

	s.server.GET("api/user/:id", s.user.Get)
	s.server.GET("api/user/list", s.user.List)
	s.server.GET("api/user/statistic/:id", s.statistic.GetUserMapStatistics)

	s.server.GET("api/log/list", s.logs.List)

	s.server.GET("api/following/list", s.following.List)
	s.server.POST("api/following/create/:code", s.following.Create)

	s.server.GET("api/beatmapset/:id", s.mapset.Get)
	s.server.GET("api/beatmapset/list", s.mapset.List)
	s.server.GET("api/beatmapset/list_for_user/:id", s.mapset.ListForUser)

	s.server.GET("api/system/statistic", s.statistic.GetSystemStatistics)
}
