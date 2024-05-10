package osuapi

import (
	"context"
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

	Interface interface {
		GetUser(ctx context.Context, userID string) (*User, error)
		GetUserMapsets(ctx context.Context, userID string) ([]*Mapset, error)
		GetUserWithMapsets(ctx context.Context, userID string) (*User, []*MapsetExtended, error)
		GetMapsetExtended(ctx context.Context, mapsetID string) (*MapsetLangGenre, error)
		GetOutgoingRequestCount() int
		ResetOutgoingRequestCount()
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
