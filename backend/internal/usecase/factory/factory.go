package factory

import (
	"osu-dashboard/internal/config"
	"osu-dashboard/internal/database/repository/beatmap"
	repositoryfactory "osu-dashboard/internal/database/repository/factory"
	"osu-dashboard/internal/database/repository/following"
	jobrepository "osu-dashboard/internal/database/repository/job"
	"osu-dashboard/internal/database/repository/log"
	"osu-dashboard/internal/database/repository/mapset"
	"osu-dashboard/internal/database/repository/user"
	"osu-dashboard/internal/database/txmanager"
	"osu-dashboard/internal/service/osuapi"
	"osu-dashboard/internal/usecase/cleanstats"
	"osu-dashboard/internal/usecase/cleanusers"
	enricherusecase "osu-dashboard/internal/usecase/enrich"
	followingcreate "osu-dashboard/internal/usecase/following/create"
	followingprovide "osu-dashboard/internal/usecase/following/provide"
	jobrecordusecase "osu-dashboard/internal/usecase/jobrecord"
	logcreate "osu-dashboard/internal/usecase/log/create"
	logprovide "osu-dashboard/internal/usecase/log/provide"
	mapsetcreate "osu-dashboard/internal/usecase/mapset/create"
	mapsetprovide "osu-dashboard/internal/usecase/mapset/provide"
	searchusecase "osu-dashboard/internal/usecase/search"
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
	LogRepo       logrepository.Interface
	JobRepo       jobrepository.Interface
}

func New(
	cfg *config.Config,
	lg *log.Logger,
	txManager txmanager.TxManager,
	osuApi *osuapi.Service,
	repoFactory *repositoryfactory.Factory,
) *UseCaseFactory {
	repos := &Repositories{
		UserRepo:      repoFactory.NewUserRepository(),
		BeatmapRepo:   repoFactory.NewBeatmapRepository(),
		MapsetRepo:    repoFactory.NewMapsetRepository(),
		FollowingRepo: repoFactory.NewFollowingsRepository(),
		LogRepo:       repoFactory.NewLogsRepository(),
		JobRepo:       repoFactory.NewJobRepository(),
	}

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
		f.repos.LogRepo,
	)
}

func (f *UseCaseFactory) MakeCreateFollowingUseCase() *followingcreate.UseCase {
	return followingcreate.New(
		f.cfg,
		f.lg,
		f.txManager,
		f.repos.FollowingRepo,
		f.MakeTrackUseCase(),
		f.repos.LogRepo,
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
	)
}

func (f *UseCaseFactory) MakeProvideLogsUseCase() *logprovide.UseCase {
	return logprovide.New(
		f.txManager,
		f.repos.LogRepo,
	)
}

func (f *UseCaseFactory) MakeCleanStatsUseCase() *cleanstats.UseCase {
	return cleanstats.New(
		f.cfg,
		f.lg,
		f.txManager,
		f.repos.UserRepo,
		f.repos.MapsetRepo,
		f.repos.BeatmapRepo,
		f.repos.LogRepo,
	)
}

func (f *UseCaseFactory) MakeCleanUsersUseCase() *cleanusers.UseCase {
	return cleanusers.New(
		f.cfg,
		f.lg,
		f.txManager,
		f.repos.UserRepo,
		f.repos.MapsetRepo,
		f.repos.LogRepo,
	)
}

func (f *UseCaseFactory) MakeEnricherUseCase() *enricherusecase.UseCase {
	return enricherusecase.New(
		f.cfg,
		f.lg,
		f.txManager,
		f.osuApi,
		f.repos.MapsetRepo,
		f.repos.FollowingRepo,
		f.repos.LogRepo,
	)
}

func (f *UseCaseFactory) MakeSearchUseCase() *searchusecase.UseCase {
	return searchusecase.New(
		f.txManager,
		f.repos.UserRepo,
		f.repos.MapsetRepo,
	)
}

func (f *UseCaseFactory) MakeCleanStatsRecorderUseCase() *jobrecordusecase.UseCase {
	return f.makeRecorderUseCase(jobrecordusecase.JobTypeCleanStats)
}

func (f *UseCaseFactory) MakeCleanUsersRecorderUseCase() *jobrecordusecase.UseCase {
	return f.makeRecorderUseCase(jobrecordusecase.JobTypeCleanUsers)
}

func (f *UseCaseFactory) MakeTrackRecorderUseCase() *jobrecordusecase.UseCase {
	return f.makeRecorderUseCase(jobrecordusecase.JobTypeTrackUsers)
}

func (f *UseCaseFactory) MakeEnrichRecorderUseCase() *jobrecordusecase.UseCase {
	return f.makeRecorderUseCase(jobrecordusecase.JobTypeEnrichData)
}

func (f *UseCaseFactory) makeRecorderUseCase(jobType jobrecordusecase.JobType) *jobrecordusecase.UseCase {
	return jobrecordusecase.New(
		f.cfg,
		f.lg,
		jobType,
		f.txManager,
		f.repos.JobRepo,
	)
}
