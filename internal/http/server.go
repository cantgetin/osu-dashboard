package http

import (
	"context"
	"github.com/ds248a/closer"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"playcount-monitor-backend/internal/app/userserviceapi"
	"playcount-monitor-backend/internal/config"
)

type Server struct {
	cfg    *config.Config
	server *echo.Echo
	lg     *log.Logger
	user   userserviceapi.ServiceImpl
}

func New(
	cfg *config.Config, lg *log.Logger, user userserviceapi.ServiceImpl,
) (*Server, error) {
	server := echo.New()
	server.HideBanner = true
	server.HidePort = true
	server.Use(middleware.CORS())

	return &Server{
		cfg:    cfg,
		server: server,
		lg:     lg,
		user:   user,
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
