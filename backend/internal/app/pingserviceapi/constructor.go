package pingserviceapi

import (
	log "github.com/sirupsen/logrus"
)

type ServiceImpl struct {
	lg *log.Logger
}

func New(
	lg *log.Logger,
) *ServiceImpl {
	return &ServiceImpl{
		lg: lg,
	}
}
