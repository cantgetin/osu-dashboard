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
		httpClient    HTTPClientInterface
	}

	Interface interface {
		GetUser(ctx context.Context, userID string) (*User, error)
		GetUserMapsets(ctx context.Context, userID string) ([]*Mapset, error)
		GetUserWithMapsets(ctx context.Context, userID string) (*User, []*Mapset, error)
	}

	HTTPClientInterface interface {
		Do(req *http.Request) (*http.Response, error)
	}
)

func New(
	cfg *config.Config,
	tokenProvider osuapitokenprovider.Interface,
	httpClient HTTPClientInterface,
) *Service {
	return &Service{
		cfg:           cfg,
		tokenProvider: tokenProvider,
		httpClient:    httpClient,
	}
}
