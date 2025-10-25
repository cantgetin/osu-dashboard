package osuapitokenprovider

import (
	"context"
	"playcount-monitor-backend/internal/bootstrap"
	"playcount-monitor-backend/internal/config"
	"sync"
	"time"
)

type (
	TokenProvider struct {
		cfg        *config.Config
		httpClient *bootstrap.CustomHTTPClient

		token      string
		validUntil time.Time
		mu         sync.Mutex
	}

	Interface interface {
		GetToken(ctx context.Context) (string, error)
	}
)

func New(cfg *config.Config, httpClient *bootstrap.CustomHTTPClient) *TokenProvider {
	return &TokenProvider{
		cfg:        cfg,
		httpClient: httpClient,
	}
}
