package followingserviceapi

import (
	"context"
	log "github.com/sirupsen/logrus"
	"osu-dashboard/internal/dto"
)

type (
	followingCreator interface {
		Create(ctx context.Context, code string) error
	}

	followingProvider interface {
		List(ctx context.Context) ([]*dto.Following, error)
	}
)

type Handlers struct {
	lg                *log.Logger
	followingCreator  followingCreator
	followingProvider followingProvider
}

func New(
	lg *log.Logger,
	followingCreator followingCreator,
	followingProvider followingProvider,
) *Handlers {
	return &Handlers{
		lg:                lg,
		followingCreator:  followingCreator,
		followingProvider: followingProvider,
	}
}
