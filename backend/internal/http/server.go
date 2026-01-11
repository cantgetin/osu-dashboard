package http

import (
	"context"
	"osu-dashboard/internal/app/http/followinghandlers"
	"osu-dashboard/internal/app/http/loghandlers"
	"osu-dashboard/internal/app/http/mapsethandlers"
	"osu-dashboard/internal/app/http/pinghandlers"
	"osu-dashboard/internal/app/http/searchhandlers"
	"osu-dashboard/internal/app/http/statistichandlers"
	"osu-dashboard/internal/app/http/usercardhandlers"
	"osu-dashboard/internal/app/http/userhandlers"
	"osu-dashboard/internal/config"
	"osu-dashboard/internal/usecase/factory"

	"github.com/ds248a/closer"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	cfg       *config.Config
	server    *echo.Echo
	lg        *log.Logger
	user      *userhandlers.Handlers
	ping      *pinghandlers.Handlers
	userCard  *usercardhandlers.Handlers
	following *followinghandlers.Handlers
	mapset    *mapsethandlers.Handlers
	statistic *statistichandlers.Handlers
	logs      *loghandlers.Handlers
	search    *searchhandlers.Handlers
}

func New(cfg *config.Config, lg *log.Logger, f *factory.UseCaseFactory) *Server {
	server := echo.New()
	server.HideBanner = true
	server.HidePort = true
	server.Use(middleware.CORS())

	ping := pinghandlers.New(lg)

	user := userhandlers.New(
		lg,
		f.MakeCreateUserUseCase(),
		f.MakeProvideUserUseCase(),
		f.MakeUpdateUserUseCase(),
	)

	userCard := usercardhandlers.New(
		lg,
		f.MakeCreateUserCardUseCase(),
		f.MakeProvideUserCardUseCase(),
		f.MakeUpdateUserCardUseCase(),
	)

	following := followinghandlers.New(
		lg,
		f.MakeCreateFollowingUseCase(),
		f.MakeProvideFollowingUseCase(),
	)

	mapset := mapsethandlers.New(
		lg,
		f.MakeProvideMapsetUseCase(),
		f.MakeCreateMapsetUseCase(),
	)

	statistic := statistichandlers.New(
		lg,
		f.MakeProvideStatisticUseCase(),
	)

	logs := loghandlers.New(
		lg,
		f.MakeProvideLogsUseCase(),
	)

	search := searchhandlers.New(lg, f.MakeSearchUseCase())

	return &Server{
		cfg:       cfg,
		server:    server,
		lg:        lg,
		ping:      ping,
		user:      user,
		userCard:  userCard,
		following: following,
		mapset:    mapset,
		statistic: statistic,
		logs:      logs,
		search:    search,
	}
}

func (s *Server) Start() {
	s.setupRoutes()

	go func() {
		s.lg.Printf("starting listening http srv at %s", s.cfg.HTTPAddr)
		s.server.Logger.Fatal(s.server.Start(s.cfg.HTTPAddr))
	}()

	closer.Add(func() {
		_ = s.Close()
	})
}

func (s *Server) Close() error {
	if err := s.server.Shutdown(context.Background()); err != nil {
		s.lg.Fatalf("error stop http srv, err: %+v", err)
		return err
	}

	s.lg.Print("http server shutdown done")

	return nil
}
