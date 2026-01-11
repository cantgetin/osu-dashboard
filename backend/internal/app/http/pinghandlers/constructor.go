package pinghandlers

import (
	log "github.com/sirupsen/logrus"
)

type Handlers struct {
	lg *log.Logger
}

func New(lg *log.Logger) *Handlers {
	return &Handlers{
		lg: lg,
	}
}
