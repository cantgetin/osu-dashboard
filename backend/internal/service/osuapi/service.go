package osuapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type BeatmapType string

const (
	Graveyard BeatmapType = "graveyard"
	Loved     BeatmapType = "loved"
	Nominated BeatmapType = "nominated"
	Pending   BeatmapType = "pending"
	Ranked    BeatmapType = "ranked"
)

func (s *Service) GetUser(ctx context.Context, userID string) (*User, error) {
	token, err := s.tokenProvider.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	// https://osu.ppy.sh/api/v2/users/123/osu
	req, err := http.NewRequest("GET", s.cfg.OsuAPIHost+"/users/"+userID+"/osu", nil)
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

	// Parsing JSON response
	var user *User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return user, nil
}

func (s *Service) GetUserMapsets(ctx context.Context, userID string) ([]*Mapset, error) {

	var BeatmapTypes = []BeatmapType{Graveyard, Loved, Nominated, Pending, Ranked}

	token, err := s.tokenProvider.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Accept":        "application/json",
		"Authorization": "Bearer " + token,
	}

	beatmapsets := []*Mapset{}
	for _, beatmapType := range BeatmapTypes {
		req, err := http.NewRequest("GET", s.cfg.OsuAPIHost+"/users/"+userID+"beatmapsets/"+string(beatmapType), nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create http request for user %s and beatmap type %s: %w", userID, beatmapType, err)
		}
		for key, value := range headers {
			req.Header.Set(key, value)
		}

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("failed to invoke request to %s: %w", req.URL.String(), err)
		}

		var maps []*Mapset
		err = json.NewDecoder(res.Body).Decode(&maps)
		if err != nil {
			return nil, fmt.Errorf("failed to decode response body: %w", err)
		}

		beatmapsets = append(beatmapsets, maps...)
		res.Body.Close()
	}

	return beatmapsets, nil
}
