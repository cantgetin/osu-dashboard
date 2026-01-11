package http

func (s *Server) setupRoutes() {
	s.server.GET("api/ping", s.handlers.ping.Ping)

	s.server.GET("api/user/:id", s.handlers.user.Get)
	s.server.GET("api/user/list", s.handlers.user.List)
	s.server.GET("api/user/statistic/:id", s.handlers.statistic.GetUserMapStatistics)

	s.server.GET("api/log/list", s.handlers.logs.List)

	s.server.GET("api/following/list", s.handlers.following.List)
	s.server.POST("api/following/create/:code", s.handlers.following.Create)

	s.server.GET("api/beatmapset/:id", s.handlers.mapset.Get)
	s.server.GET("api/beatmapset/list", s.handlers.mapset.List)
	s.server.GET("api/beatmapset/list_for_user/:id", s.handlers.mapset.ListForUser)

	s.server.GET("api/system/statistic", s.handlers.statistic.GetSystemStatistics)

	s.server.POST("/api/search/:query", s.handlers.search.Search)
}
