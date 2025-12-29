package osuapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"golang.org/x/sync/errgroup"
)

func (s *Service) GetUserMapsets(ctx context.Context, userID string) ([]*Mapset, error) {
	var mapsetTypes = []MapsetStatusAPIOption{Graveyard, Loved, Pending, Ranked}

	token, err := s.tokenProvider.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Accept":        "application/json",
		"Authorization": "Bearer " + token,
	}

	var (
		mu             sync.Mutex
		allBeatmapsets []*Mapset
	)

	eg, egCtx := errgroup.WithContext(ctx)

	for _, mapsetType := range mapsetTypes {
		eg.Go(func() error {
			var beatmapsets []*Mapset
			beatmapsets, err = s.fetchBeatmapsets(egCtx, userID, string(mapsetType), 0, headers)
			if err != nil {
				return fmt.Errorf("failed to fetch %s beatmapsets: %w", mapsetType, err)
			}

			mu.Lock()
			allBeatmapsets = append(allBeatmapsets, beatmapsets...)
			mu.Unlock()

			return nil
		})
	}

	if err = eg.Wait(); err != nil {
		return nil, err
	}

	return allBeatmapsets, nil
}

func (s *Service) fetchBeatmapsets(
	ctx context.Context,
	userID string,
	mapsetType string,
	offset int,
	headers map[string]string,
) ([]*Mapset, error) {
	fetchURL := s.cfg.OsuAPIHost + "/users/" + userID + "/beatmapsets/" + mapsetType + "?limit=100&offset=" + strconv.Itoa(offset)
	req, err := http.NewRequestWithContext(ctx, "GET", fetchURL, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request for user %s and beatmap type %s: %w", userID, mapsetType, err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	res, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to invoke request to %s: %w", req.URL.String(), err)
	}
	defer res.Body.Close()

	var maps []*Mapset
	err = json.NewDecoder(res.Body).Decode(&maps)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	if len(maps) >= 100 {
		nextPageMaps, err := s.fetchBeatmapsets(ctx, userID, mapsetType, offset+100, headers)
		if err != nil {
			return nil, err
		}
		maps = append(maps, nextPageMaps...)
	}

	return maps, nil
}

func (s *Service) GetMapsetExtended(ctx context.Context, mapsetID string) (*MapsetLangGenre, error) {
	token, err := s.tokenProvider.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	// https://osu.ppy.sh/api/v2/beatmapsets
	req, err := http.NewRequestWithContext(ctx, "GET", s.cfg.OsuAPIHost+"/beatmapsets/"+mapsetID, http.NoBody)
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

	var mapsetExt *MapsetLangGenre
	err = json.NewDecoder(resp.Body).Decode(&mapsetExt)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return mapsetExt, nil
}

func (s *Service) GetMapsetCommentsCount(ctx context.Context, mapsetID string) (int, error) {
	token, err := s.tokenProvider.GetToken(ctx)
	if err != nil {
		return 0, err
	}

	fetchURL := s.cfg.OsuAPIHost + "/comments?commentable_type=beatmapset&commentable_id=" + mapsetID + "&sort=new"

	// https://osu.ppy.sh/api/v2/comments
	req, err := http.NewRequestWithContext(ctx, "GET", fetchURL, http.NoBody)
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
