package userrepository

import (
	"osu-dashboard/internal/config"

	log "github.com/sirupsen/logrus"
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
