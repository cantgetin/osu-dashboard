package osuapi

import (
	"net/http"
	"playcount-monitor-backend/internal/config"
	"playcount-monitor-backend/internal/service/osuapitokenprovider"
)

type (
	Service struct {
		cfg           *config.Config
		tokenProvider osuapitokenprovider.Interface
		httpClient    *http.Client
	}
)

func New(
	cfg *config.Config,
	tokenProvider osuapitokenprovider.Interface,
	httpClient *http.Client,
) *Service {
	return &Service{
		cfg:           cfg,
		tokenProvider: tokenProvider,
		httpClient:    httpClient,
	}
}
