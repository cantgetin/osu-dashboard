package osuapitokenprovider

import (
	"context"
	"net/http"
	"playcount-monitor-backend/internal/config"
	"sync"
	"time"
)

type (
	TokenProvider struct {
		cfg        *config.Config
		httpClient *http.Client

		token      string
		validUntil time.Time
		mu         sync.Mutex
	}

	Interface interface {
		GetToken(ctx context.Context) (string, error)
	}
)

func New(cfg *config.Config, httpClient *http.Client) *TokenProvider {
	return &TokenProvider{
		cfg:        cfg,
		httpClient: httpClient,
	}
}
