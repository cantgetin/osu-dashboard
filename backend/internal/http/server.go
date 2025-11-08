package http

import (
	"context"
	"github.com/ds248a/closer"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"osu-dashboard/internal/app/followingserviceapi"
	"osu-dashboard/internal/app/logserviceapi"
	"osu-dashboard/internal/app/mapsetserviceapi"
	"osu-dashboard/internal/app/pingserviceapi"
	"osu-dashboard/internal/app/statisticserviceapi"
	"osu-dashboard/internal/app/usercardserviseapi"
	"osu-dashboard/internal/app/userserviceapi"
	"osu-dashboard/internal/config"
	"osu-dashboard/internal/usecase/factory"
)

type Server struct {
	cfg       *config.Config
	server    *echo.Echo
	lg        *log.Logger
	user      *userserviceapi.ServiceImpl
	ping      *pingserviceapi.ServiceImpl
	userCard  *usercardserviseapi.ServiceImpl
	following *followingserviceapi.ServiceImpl
	mapset    *mapsetserviceapi.ServiceImpl
	statistic *statisticserviceapi.ServiceImpl
	logs      *logserviceapi.ServiceImpl
}

func New(
	cfg *config.Config, lg *log.Logger, f *factory.UseCaseFactory,
) (*Server, error) {
	server := echo.New()
	server.HideBanner = true
	server.HidePort = true
	server.Use(middleware.CORS())

	ping := pingserviceapi.New(lg)

	user := userserviceapi.New(
		lg,
		f.MakeCreateUserUseCase(),
		f.MakeProvideUserUseCase(),
		f.MakeUpdateUserUseCase(),
	)

	userCard := usercardserviseapi.New(
		lg,
		f.MakeCreateUserCardUseCase(),
		f.MakeProvideUserCardUseCase(),
		f.MakeUpdateUserCardUseCase(),
	)

	following := followingserviceapi.New(
		lg,
		f.MakeCreateFollowingUseCase(),
		f.MakeProvideFollowingUseCase(),
	)

	mapset := mapsetserviceapi.New(
		lg,
		f.MakeProvideMapsetUseCase(),
		f.MakeCreateMapsetUseCase(),
	)

	statistic := statisticserviceapi.New(
		lg,
		f.MakeProvideStatisticUseCase(),
	)

	logs := logserviceapi.New(
		lg,
		f.MakeProvideLogsUseCase(),
	)

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
	}, nil
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
