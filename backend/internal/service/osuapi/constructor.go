package osuapi

import (
	"playcount-monitor-backend/internal/bootstrap"
	"playcount-monitor-backend/internal/config"
	"playcount-monitor-backend/internal/service/osuapitokenprovider"
)

type (
	Service struct {
		cfg           *config.Config
		tokenProvider osuapitokenprovider.Interface
		httpClient    *bootstrap.CustomHTTPClient
	}
)

func New(
	cfg *config.Config,
	tokenProvider osuapitokenprovider.Interface,
	httpClient *bootstrap.CustomHTTPClient,
) *Service {
	return &Service{
		cfg:           cfg,
		tokenProvider: tokenProvider,
		httpClient:    httpClient,
	}
}
