package followinghandlers

import (
	"context"
	"osu-dashboard/internal/dto"

	log "github.com/sirupsen/logrus"
)

type (
	followingCreator interface {
		Create(ctx context.Context, code string) error
	}

	followingProvider interface {
		List(ctx context.Context) ([]*dto.Following, error)
	}

	Handlers struct {
		lg                *log.Logger
		followingCreator  followingCreator
		followingProvider followingProvider
	}
)

func New(lg *log.Logger, c followingCreator, p followingProvider) *Handlers {
	return &Handlers{
		lg:                lg,
		followingCreator:  c,
		followingProvider: p,
	}
}
