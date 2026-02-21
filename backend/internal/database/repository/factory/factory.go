package repositoryfactory

import (
	"osu-dashboard/internal/config"
	"osu-dashboard/internal/database/repository/beatmap"
	"osu-dashboard/internal/database/repository/following"
	"osu-dashboard/internal/database/repository/job"
	"osu-dashboard/internal/database/repository/log"
	"osu-dashboard/internal/database/repository/mapset"
	"osu-dashboard/internal/database/repository/user"

	log "github.com/sirupsen/logrus"
)

type Factory struct {
	lg  *log.Logger
	cfg *config.Config
}

func New(cfg *config.Config, lg *log.Logger) *Factory {
	return &Factory{
		lg:  lg,
		cfg: cfg,
	}
}

func (f *Factory) NewBeatmapRepository() *beatmaprepository.GormRepository {
	return beatmaprepository.New(f.cfg, f.lg)
}

func (f *Factory) NewFollowingsRepository() *followingrepository.GormRepository {
	return followingrepository.New(f.cfg, f.lg)
}

func (f *Factory) NewLogsRepository() *logrepository.GormRepository {
	return logrepository.New(f.cfg, f.lg)
}

func (f *Factory) NewMapsetRepository() *mapsetrepository.GormRepository {
	return mapsetrepository.New(f.cfg, f.lg)
}

func (f *Factory) NewUserRepository() *userrepository.GormRepository {
	return userrepository.New(f.cfg, f.lg)
}

func (f *Factory) NewJobRepository() *jobrepository.GormRepository {
	return jobrepository.New(f.cfg, f.lg)
}
