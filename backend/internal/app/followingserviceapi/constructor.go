package followingserviceapi

import (
	"context"
	log "github.com/sirupsen/logrus"
	"osu-dashboard/internal/dto"
)

type followingCreator interface {
	Create(ctx context.Context, code string) error
}

type followingProvider interface {
	List(ctx context.Context) ([]*dto.Following, error)
}

type ServiceImpl struct {
	lg                *log.Logger
	followingCreator  followingCreator
	followingProvider followingProvider
}

func New(
	lg *log.Logger,
	followingCreator followingCreator,
	followingProvider followingProvider,
) *ServiceImpl {
	return &ServiceImpl{
		lg:                lg,
		followingCreator:  followingCreator,
		followingProvider: followingProvider,
	}
}
