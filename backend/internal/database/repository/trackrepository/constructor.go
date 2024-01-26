package trackrepository

import (
	log "github.com/sirupsen/logrus"
	"playcount-monitor-backend/internal/config"
)

type GormRepository struct {
	lg  *log.Logger
	cfg *config.Config
}

func New(cfg *config.Config, lg *log.Logger) (*GormRepository, error) {
	return &GormRepository{
		lg:  lg,
		cfg: cfg,
	}, nil
}
