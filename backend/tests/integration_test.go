package tests

import (
	"context"
	"github.com/caarlos0/env"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"playcount-monitor-backend/internal/app"
	"playcount-monitor-backend/internal/config"
	"playcount-monitor-backend/tests/integration"
	"testing"
	"time"
)

func TestIntegration(t *testing.T) {
	suite.Run(t, &IntegrationSuite{})
}

type IntegrationSuite struct {
	suite.Suite
	cfg       *config.Config
	ctx       context.Context
	cancelCtx func()
	closers   []func() error
	db        *gorm.DB
}

func (s *IntegrationSuite) SetupSuite() {
	s.cfg = &config.Config{}
	if err := env.Parse(s.cfg); err != nil {
		s.T().Fatalf("failed to parse cfg, %v", err)
	}

	if !s.cfg.RunIntegrationTest {
		s.T().SkipNow()
	}

	s.ctx, s.cancelCtx = context.WithCancel(
		context.WithValue(context.Background(), "environment", "integration-test"),
	)

	s.T().Log("Starting Docker containers...")
	pool, dockerClose := integration.Start(s.T(), s.cfg)
	s.closers = append(s.closers, dockerClose)

	s.T().Log("Initializing DB with migrations...")
	gdb, closeDB := integration.InitDB(s.T(), pool, s.cfg)
	s.closers = append(s.closers, closeDB)
	s.db = gdb

	s.T().Log("Setup completed")

	lg := log.New()

	go func() {
		if err := app.Run(s.ctx, s.cfg, lg); err != nil {
			s.T().Logf("application has exited %v", err)
		}
	}()

	time.Sleep(5 * time.Second)
}

func (s *IntegrationSuite) TearDownSuite() {
	s.T().Log("Suite teardown...")
	s.cancelCtx()
	s.closeAll()
}

func (s *IntegrationSuite) closeAll() {
	for _, c := range s.closers {
		if err := c(); err != nil {
			s.Error(err)
		}
	}
}
