package factory

import (
	log "github.com/sirupsen/logrus"
	"playcount-monitor-backend/internal/config"
	"playcount-monitor-backend/internal/database/repository/beatmaprepository"
	"playcount-monitor-backend/internal/database/repository/mapsetrepository"
	"playcount-monitor-backend/internal/database/repository/userrepository"
	"playcount-monitor-backend/internal/database/txmanager"
	mapsetcreate "playcount-monitor-backend/internal/usecase/mapset/create"
	usercreate "playcount-monitor-backend/internal/usecase/user/create"
	userprovide "playcount-monitor-backend/internal/usecase/user/provide"
	userupdate "playcount-monitor-backend/internal/usecase/user/update"
)

type UseCaseFactory struct {
	lg        *log.Logger
	cfg       *config.Config
	txManager txmanager.TxManager
	repos     *Repositories
}

type Repositories struct {
	UserRepo    userrepository.Interface
	BeatmapRepo beatmaprepository.Interface
	MapsetRepo  mapsetrepository.Interface
}

func New(
	cfg *config.Config,
	lg *log.Logger,
	txManager txmanager.TxManager,
	repos *Repositories,
) (*UseCaseFactory, error) {
	return &UseCaseFactory{
		cfg:       cfg,
		lg:        lg,
		txManager: txManager,
		repos:     repos,
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
