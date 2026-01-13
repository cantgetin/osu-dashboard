package osuapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	osuapimodels "osu-dashboard/internal/service/osuapi/models"
	"strings"
)

func (s *Service) GetUserInfoByHisToken(ctx context.Context, accessToken string) (*osuapimodels.User, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://osu.ppy.sh/api/v2/me", http.NoBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Accept", "application/json")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make http request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("api server returned %d: %s", resp.StatusCode, body)
	}

	var user osuapimodels.User
	if err = json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

// ExchangeCodeForToken using this method instead of tokenprovider cause it exchanges user code for token
func (s *Service) ExchangeCodeForToken(ctx context.Context, code string) (*osuapimodels.TokenResponse, error) {
	data := url.Values{}
	data.Set("client_id", s.cfg.OsuAPIClientID)
	data.Set("client_secret", s.cfg.OsuAPIClientSecret)
	data.Set("code", code)
	data.Set("grant_type", "authorization_code")
	data.Set("redirect_uri", s.cfg.OsuAPIRedirectURI)

	req, err := http.NewRequestWithContext(ctx, "POST", "https://osu.ppy.sh/oauth/token", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make http request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("oauth server returned %d: %s", resp.StatusCode, body)
	}

	var token osuapimodels.TokenResponse
	if err = json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return nil, err
	}

	return &token, nil
}
