package osuapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type BeatmapType string

const (
	Graveyard BeatmapType = "graveyard"
	Loved     BeatmapType = "loved"
	// Nominated we not use this cause it shows maps that user nominated (from others) which breaks mapset FK
	Nominated BeatmapType = "nominated"
	Pending   BeatmapType = "pending"
	Ranked    BeatmapType = "ranked"
)

func (s *Service) GetMapsetCommentsCount(ctx context.Context, mapsetID string) (int, error) {
	token, err := s.tokenProvider.GetToken(ctx)
	if err != nil {
		return 0, err
	}

	// https://osu.ppy.sh/api/v2/comments
	req, err := http.NewRequest("GET",
		s.cfg.OsuAPIHost+"/comments?commentable_type=beatmapset&commentable_id="+mapsetID+"&sort=new", nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create http request: %w", err)
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
		return 0, fmt.Errorf("failed to invoke request: %w", err)
	}
	defer resp.Body.Close()

	// Parsing JSON response
	var comments *Comments
	err = json.NewDecoder(resp.Body).Decode(&comments)
	if err != nil {
		return 0, fmt.Errorf("failed to decode response body: %w", err)
	}

	return comments.Total, nil
}

func (s *Service) GetUserWithMapsets(ctx context.Context, userID string) (*User, []*Mapset, error) {
	user, err := s.GetUser(ctx, userID)
	if err != nil {
		return nil, nil, err
	}

	userMapsets, err := s.GetUserMapsets(ctx, userID)
	if err != nil {
		return nil, nil, err
	}

	for _, mapset := range userMapsets {
		mapset.CommentsCount, err = s.GetMapsetCommentsCount(ctx, strconv.Itoa(mapset.Id))
		if err != nil {
			return nil, nil, err
		}
	}

	return user, userMapsets, nil
}

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
	var BeatmapTypes = []BeatmapType{Graveyard, Loved, Pending, Ranked}

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
		beatmapsets, err = s.fetchBeatmapsets(userID, string(beatmapType), 0, headers, beatmapsets)
		if err != nil {
			return nil, err
		}
	}

	return beatmapsets, nil
}

func (s *Service) fetchBeatmapsets(
	userID string,
	beatmapType string,
	offset int,
	headers map[string]string,
	beatmapsets []*Mapset,
) ([]*Mapset, error) {
	req, err := http.NewRequest("GET", s.cfg.OsuAPIHost+"/users/"+userID+"/beatmapsets/"+beatmapType+"?limit=100&offset="+strconv.Itoa(offset), nil)
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

	if len(maps) >= 100 {
		// If there are 100 or more maps, fetch the next page
		return s.fetchBeatmapsets(userID, beatmapType, offset+100, headers, beatmapsets)
	}

	return beatmapsets, nil
}
