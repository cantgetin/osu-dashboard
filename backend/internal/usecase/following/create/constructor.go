package followingcreate

import (
	"context"
	"playcount-monitor-backend/internal/config"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
	"playcount-monitor-backend/internal/service/osuapi"

	log "github.com/sirupsen/logrus"
)

type followingSource interface {
	Create(ctx context.Context, tx txmanager.Tx, user *model.Following) error
	Get(ctx context.Context, tx txmanager.Tx, id int) (*model.Following, error)
}

type logSource interface {
	Create(ctx context.Context, log *model.Log) error
}

type trackSource interface {
	TrackSingleFollowing(ctx context.Context, following *model.Following) error
}

type osuAPI interface {
	GetUserInfoByHisToken(accessToken string) (*osuapi.UserResponse, error)
	ExchangeCodeForToken(code string) (*osuapi.TokenResponse, error)
	GetTransportStats() osuapi.TransportStats
	ResetStats()
}

type UseCase struct {
	cfg       *config.Config
	lg        *log.Logger
	txm       txmanager.TxManager
	following followingSource
	log       logSource
	track     trackSource
	osuAPI    osuAPI
}

func New(
	cfg *config.Config,
	lg *log.Logger,
	txm txmanager.TxManager,
	following followingSource,
	track trackSource,
	log logSource,
	osuAPI osuAPI,
) *UseCase {
	return &UseCase{
		cfg:       cfg,
		lg:        lg,
		txm:       txm,
		following: following,
		log:       log,
		track:     track,
		osuAPI:    osuAPI,
	}
}
