package factory

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
	"osu-dashboard/internal/database/txmanager"
	"osu-dashboard/internal/service/osuapi"
	cleanerUseCase "osu-dashboard/internal/usecase/cleaner"
	enricherusecase "osu-dashboard/internal/usecase/enricher"
	followingcreate "osu-dashboard/internal/usecase/following/create"
	followingprovide "osu-dashboard/internal/usecase/following/provide"
	logcreate "osu-dashboard/internal/usecase/log/create"
	logprovide "osu-dashboard/internal/usecase/log/provide"
	mapsetcreate "osu-dashboard/internal/usecase/mapset/create"
	mapsetprovide "osu-dashboard/internal/usecase/mapset/provide"
	statisticprovide "osu-dashboard/internal/usecase/statistic/provide"
	"osu-dashboard/internal/usecase/track"
	usercreate "osu-dashboard/internal/usecase/user/create"
	userprovide "osu-dashboard/internal/usecase/user/provide"
	userupdate "osu-dashboard/internal/usecase/user/update"
	usercardcreate "osu-dashboard/internal/usecase/usercard/create"
	usercardprovide "osu-dashboard/internal/usecase/usercard/provide"
	usercardupdate "osu-dashboard/internal/usecase/usercard/update"

	log "github.com/sirupsen/logrus"
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
	EnrichesRepo  enrichesrepository.Interface
	TrackRepo     trackrepository.Interface
	LogRepo       logrepository.Interface
	CleanRepo     cleanrepository.Interface
}

func New(
	cfg *config.Config,
	lg *log.Logger,
	txManager txmanager.TxManager,
	osuApi *osuapi.Service,
	repos *Repositories,
) *UseCaseFactory {
	return &UseCaseFactory{
		cfg:       cfg,
		lg:        lg,
		txManager: txManager,
		repos:     repos,
		osuApi:    osuApi,
	}
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
		f.lg,
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

func (f *UseCaseFactory) MakeCleanerUseCase() *cleanerUseCase.UseCase {
	return cleanerUseCase.New(
		f.cfg,
		f.lg,
		f.txManager,
		f.repos.UserRepo,
		f.repos.MapsetRepo,
		f.repos.BeatmapRepo,
		f.MakeCreateLogUseCase(),
		f.repos.CleanRepo)
}

func (f *UseCaseFactory) MakeEnricherUseCase() *enricherusecase.UseCase {
	return enricherusecase.New(
		f.cfg,
		f.lg,
		f.txManager,
		f.osuApi,
		f.repos.MapsetRepo,
		f.repos.FollowingRepo,
		f.repos.EnrichesRepo,
	)
}
