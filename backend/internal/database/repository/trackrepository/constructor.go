package trackrepository

import (
	log "github.com/sirupsen/logrus"
	"osu-dashboard/internal/config"
)

type GormRepository struct {
	lg  *log.Logger
	cfg *config.Config
}

func New(cfg *config.Config, lg *log.Logger) *GormRepository {
	return &GormRepository{
		lg:  lg,
		cfg: cfg,
	}
}
