package osuapi

import (
	"osu-dashboard/internal/bootstrap"
	"osu-dashboard/internal/config"
	"osu-dashboard/internal/service/osuapitokenprovider"
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
