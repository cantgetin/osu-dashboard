package osuapi

import (
	"context"
	"net/http"
	"playcount-monitor-backend/internal/config"
)

type (
	Service struct {
		cfg        *config.Config
		httpClient HTTPClientInterface
	}

	Interface interface {
		GetUser(ctx context.Context, serviceID, token string) (*User, error)
	}

	HTTPClientInterface interface {
		Do(req *http.Request) (*http.Response, error)
	}
)
