package searchhandlers

import (
	"context"
	log "github.com/sirupsen/logrus"
	"osu-dashboard/internal/dto"
)

type (
	searcher interface {
		Search(ctx context.Context, query string) (result []*dto.SearchResult, err error)
	}
	Handlers struct {
		lg       *log.Logger
		searcher searcher
	}
)

func New(lg *log.Logger, search searcher) *Handlers {
	return &Handlers{
		lg:       lg,
		searcher: search,
	}
}
