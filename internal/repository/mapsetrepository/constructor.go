package mapsetrepository

import (
	"gorm.io/gorm"
	"log"
	"playcount-monitor-backend/internal/config"
)

type GormRepository struct {
	lg  *log.Logger
	cfg *config.Config
	db  *gorm.DB
}

func New(cfg *config.Config, lg *log.Logger, db *gorm.DB) (*GormRepository, error) {
	return &GormRepository{
		lg:  lg,
		cfg: cfg,
		db:  db,
	}, nil
}
