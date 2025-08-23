package factory

import (
	log "github.com/sirupsen/logrus"
	"playcount-monitor-backend/internal/config"
	"playcount-monitor-backend/internal/database/repository/beatmaprepository"
	"playcount-monitor-backend/internal/database/repository/followingrepository"
	"playcount-monitor-backend/internal/database/repository/logrepository"
	"playcount-monitor-backend/internal/database/repository/mapsetrepository"
	"playcount-monitor-backend/internal/database/repository/trackrepository"
	"playcount-monitor-backend/internal/database/repository/userrepository"
	"playcount-monitor-backend/internal/database/txmanager"
	"playcount-monitor-backend/internal/service/osuapi"
	followingcreate "playcount-monitor-backend/internal/usecase/following/create"
	followingprovide "playcount-monitor-backend/internal/usecase/following/provide"
	logcreate "playcount-monitor-backend/internal/usecase/log/create"
	logprovide "playcount-monitor-backend/internal/usecase/log/provide"
	mapsetcreate "playcount-monitor-backend/internal/usecase/mapset/create"
	mapsetprovide "playcount-monitor-backend/internal/usecase/mapset/provide"
	statisticprovide "playcount-monitor-backend/internal/usecase/statistic/provide"
	"playcount-monitor-backend/internal/usecase/track"
	usercreate "playcount-monitor-backend/internal/usecase/user/create"
	userprovide "playcount-monitor-backend/internal/usecase/user/provide"
	userupdate "playcount-monitor-backend/internal/usecase/user/update"
	usercardcreate "playcount-monitor-backend/internal/usecase/usercard/create"
	usercardprovide "playcount-monitor-backend/internal/usecase/usercard/provide"
	usercardupdate "playcount-monitor-backend/internal/usecase/usercard/update"
)

type UseCaseFactory struct {
	lg        *log.Logger
	cfg       *config.Config
	txManager txmanager.TxManager
	osuApi    *osuapi.Service
	repos     *Repositories
}

type Repositories struct {
	UserRepo      userrepository.Interface
	BeatmapRepo   beatmaprepository.Interface
	MapsetRepo    mapsetrepository.Interface
	FollowingRepo followingrepository.Interface
	TrackRepo     trackrepository.Interface
	LogRepo       logrepository.Interface
}

func New(
	cfg *config.Config,
	lg *log.Logger,
	txManager txmanager.TxManager,
	osuApi *osuapi.Service,
	repos *Repositories,
) (*UseCaseFactory, error) {
	return &UseCaseFactory{
		cfg:       cfg,
		lg:        lg,
		txManager: txManager,
		repos:     repos,
		osuApi:    osuApi,
	}, nil
}

func (f *UseCaseFactory) MakeCreateMapsetUseCase() *mapsetcreate.UseCase {
	return mapsetcreate.New(
		f.cfg,
		f.lg,
		f.txManager,
		f.repos.BeatmapRepo,
		f.repos.MapsetRepo,
	)
}

func (f *UseCaseFactory) MakeProvideMapsetUseCase() *mapsetprovide.UseCase {
	return mapsetprovide.New(
		f.cfg,
		f.lg,
		f.txManager,
		f.repos.BeatmapRepo,
		f.repos.MapsetRepo,
	)
}

func (f *UseCaseFactory) MakeCreateUserUseCase() *usercreate.UseCase {
	return usercreate.New(
		f.cfg,
		f.lg,
		f.txManager,
		f.repos.UserRepo,
	)
}

func (f *UseCaseFactory) MakeProvideUserUseCase() *userprovide.UseCase {
	return userprovide.New(
		f.cfg,
		f.lg,
		f.txManager,
		f.repos.UserRepo,
		f.osuApi,
	)
}

func (f *UseCaseFactory) MakeUpdateUserUseCase() *userupdate.UseCase {
	return userupdate.New(
		f.cfg,
		f.lg,
		f.txManager,
		f.repos.UserRepo,
	)
}

func (f *UseCaseFactory) MakeCreateUserCardUseCase() *usercardcreate.UseCase {
	return usercardcreate.New(
		f.cfg,
		f.lg,
		f.txManager,
		f.repos.UserRepo,
		f.repos.MapsetRepo,
		f.repos.BeatmapRepo,
	)
}

func (f *UseCaseFactory) MakeProvideUserCardUseCase() *usercardprovide.UseCase {
	return usercardprovide.New(
		f.cfg,
		f.lg,
		f.txManager,
		f.repos.UserRepo,
		f.repos.MapsetRepo,
		f.repos.BeatmapRepo,
	)
}

func (f *UseCaseFactory) MakeUpdateUserCardUseCase() *usercardupdate.UseCase {
	return usercardupdate.New(
		f.cfg,
		f.lg,
		f.txManager,
		f.repos.UserRepo,
		f.repos.MapsetRepo,
		f.repos.BeatmapRepo,
	)
}

func (f *UseCaseFactory) MakeCreateLogUseCase() *logcreate.UseCase {
	return logcreate.New(f.txManager, f.repos.LogRepo)
}

func (f *UseCaseFactory) MakeTrackUseCase() *track.UseCase {
	return track.New(
		f.cfg,
		f.txManager,
		f.osuApi,
		f.repos.UserRepo,
		f.repos.MapsetRepo,
		f.repos.BeatmapRepo,
		f.repos.FollowingRepo,
		f.repos.TrackRepo,
		f.MakeCreateLogUseCase(),
	)
}

func (f *UseCaseFactory) MakeCreateFollowingUseCase() *followingcreate.UseCase {
	return followingcreate.New(
		f.cfg,
		f.lg,
		f.txManager,
		f.repos.FollowingRepo,
		f.MakeTrackUseCase(),
		f.MakeCreateLogUseCase(),
		f.osuApi,
	)
}

func (f *UseCaseFactory) MakeProvideFollowingUseCase() *followingprovide.UseCase {
	return followingprovide.New(
		f.cfg,
		f.lg,
		f.txManager,
		f.repos.FollowingRepo,
	)
}

func (f *UseCaseFactory) MakeProvideStatisticUseCase() *statisticprovide.UseCase {
	return statisticprovide.New(
		f.cfg,
		f.lg,
		f.txManager,
		f.repos.BeatmapRepo,
		f.repos.MapsetRepo,
		f.repos.UserRepo,
		f.repos.TrackRepo,
	)
}

func (f *UseCaseFactory) MakeProvideLogsUseCase() *logprovide.UseCase {
	return logprovide.New(
		f.txManager,
		f.repos.LogRepo,
	)
}
