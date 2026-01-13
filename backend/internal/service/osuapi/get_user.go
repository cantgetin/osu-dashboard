package osuapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	osuapimodels "osu-dashboard/internal/service/osuapi/models"
)

func (s *Service) GetUser(ctx context.Context, userID string) (*osuapimodels.User, error) {
	token, err := s.tokenProvider.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	// https://osu.ppy.sh/api/v2/users/123/osu
	req, err := http.NewRequestWithContext(ctx, "GET", s.cfg.OsuAPIHost+"/users/"+userID+"/osu", http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %w", err)
	}

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Accept":        "application/json",
		"Authorization": "Bearer " + token,
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to invoke request: %w", err)
	}
	defer resp.Body.Close()

	var user *osuapimodels.User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return user, nil
}
