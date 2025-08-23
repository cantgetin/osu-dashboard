package logserviceapi

import (
	"context"
	log "github.com/sirupsen/logrus"
	"playcount-monitor-backend/internal/dto"
)

type logProvider interface {
	List(ctx context.Context, page int) (*dto.LogsPaged, error)
}

type ServiceImpl struct {
	lg          *log.Logger
	logProvider logProvider
}

func New(
	lg *log.Logger,
	logProvider logProvider,
) *ServiceImpl {
	return &ServiceImpl{
		lg:          lg,
		logProvider: logProvider,
	}
}
