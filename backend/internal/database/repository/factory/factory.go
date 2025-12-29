package repositoryfactory

import (
	"osu-dashboard/internal/config"
	"osu-dashboard/internal/database/repository/beatmaprepository"
	"osu-dashboard/internal/database/repository/cleanrepository"
	"osu-dashboard/internal/database/repository/enrichesrepository"
	"osu-dashboard/internal/database/repository/followingrepository"
	"osu-dashboard/internal/database/repository/logrepository"
	"osu-dashboard/internal/database/repository/mapsetrepository"
	"osu-dashboard/internal/database/repository/trackrepository"
	"osu-dashboard/internal/database/repository/userrepository"

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

func (f *Factory) NewCleansRepository() *cleanrepository.GormRepository {
	return cleanrepository.New(f.cfg, f.lg)
}

func (f *Factory) NewEnrichesRepository() *enrichesrepository.GormRepository {
	return enrichesrepository.New(f.cfg, f.lg)
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

func (f *Factory) NewTrackRepository() *trackrepository.GormRepository {
	return trackrepository.New(f.cfg, f.lg)
}

func (f *Factory) NewUserRepository() *userrepository.GormRepository {
	return userrepository.New(f.cfg, f.lg)
}
