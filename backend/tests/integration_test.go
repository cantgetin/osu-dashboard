package tests

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"playcount-monitor-backend/internal/app"
	"playcount-monitor-backend/internal/config"
	"playcount-monitor-backend/tests/integration"
	"testing"
	"time"
)

type ContextKey string

const EnvKey ContextKey = "environment"

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
	port      string
}

func (s *IntegrationSuite) SetupSuite() {
	var err error
	s.cfg, err = config.LoadConfig(".env_test")
	if err != nil {
		s.T().Fatalf("failed to load config, %v", err)
	}

	if !s.cfg.RunIntegrationTest {
		s.T().Log("skipping integration tests ...")
		s.T().SkipNow()
	}

	s.port = s.cfg.IntegrationTestHTTPPort
	s.cfg.HTTPAddr = s.cfg.IntegrationTestHTTPAddr

	s.ctx, s.cancelCtx = context.WithCancel(
		context.WithValue(context.Background(), EnvKey, "integration-test"),
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
