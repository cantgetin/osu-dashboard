package mappers

import (
	"encoding/json"
	"fmt"
	"playcount-monitor-backend/internal/database/repository"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/dto"
	"time"
)

func MapBeatmapDTOToBeatmapModel(beatmap *dto.Beatmap) (*model.Beatmap, error) {
	stats, err := MapBeatmapStats(beatmap)
	if err != nil {
		return nil, err
	}

	return &model.Beatmap{
		ID:               beatmap.Id,
		MapsetID:         beatmap.BeatmapsetId,
		DifficultyRating: beatmap.DifficultyRating,
		Version:          beatmap.Version,
		Accuracy:         beatmap.Accuracy,
		AR:               beatmap.Ar,
		BPM:              float64(beatmap.Bpm),
		CS:               beatmap.Cs,
		Status:           beatmap.Status,
		URL:              beatmap.Url,
		TotalLength:      beatmap.TotalLength,
		BeatmapStats:     stats,
	}, nil
}

func MapUserDTOToUserModel(user *dto.User) (*model.User, error) {
	return &model.User{
		ID:                       user.ID,
		Username:                 user.Username,
		AvatarURL:                user.AvatarURL,
		GraveyardBeatmapsetCount: user.GraveyardBeatmapsetCount,
		UnrankedBeatmapsetCount:  user.UnrankedBeatmapsetCount,
		CreatedAt:                time.Now().UTC(),
		UpdatedAt:                time.Now().UTC(),
	}, nil
}

func MapMapsetDTOToMapsetModel(mapset *dto.Mapset) (*model.Mapset, error) {
	mapsetStats, err := MapMapsetStats(mapset)
	if err != nil {
		return nil, err
	}

	covers, err := mapCovers(mapset.Covers)
	if err != nil {
		return nil, err
	}

	return &model.Mapset{
		ID:          mapset.Id,
		Artist:      mapset.Artist,
		Title:       mapset.Title,
		Covers:      covers,
		Status:      mapset.Status,
		LastUpdated: mapset.LastUpdated,
		UserID:      mapset.UserId,
		Creator:     mapset.Creator,
		PreviewURL:  mapset.PreviewUrl,
		Tags:        mapset.Tags,
		MapsetStats: mapsetStats,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}, nil
}

func MapBeatmapStats(beatmap *dto.Beatmap) (repository.JSON, error) {
	var stats = make(model.BeatmapStats)
	stats[time.Now().UTC()] = &model.BeatmapStatsModel{
		Playcount: beatmap.Playcount,
		Passcount: beatmap.Passcount,
	}

	statsJson, err := json.Marshal(stats)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal beatmap stats: %w", err)
	}

	return statsJson, nil
}

func mapCovers(m map[string]string) (repository.JSON, error) {
	coversJson, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal mapset covers to json: %w", err)
	}

	return coversJson, nil
}

func MapMapsetStats(mapset *dto.Mapset) (repository.JSON, error) {
	var stats = make(model.MapsetStats)
	stats[time.Now().UTC()] = &model.MapsetStatsModel{
		Playcount: mapset.PlayCount,
		Favorites: mapset.FavouriteCount,
	}

	statsJson, err := json.Marshal(stats)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal mapset stats to json: %w", err)
	}

	return statsJson, nil
}

func AppendNewMapsetStats(json1, json2 repository.JSON) (repository.JSON, error) {
	// merge two JSONs that are map[time.Time]model.MapsetStatsModel
	map1 := make(model.MapsetStats)
	map2 := make(model.MapsetStats)

	if err := json.Unmarshal(json1, &map1); err != nil {
		return repository.JSON{}, err
	}

	if err := json.Unmarshal(json2, &map2); err != nil {
		return repository.JSON{}, err
	}

	for key, value := range map2 {
		map1[key] = value
	}

	mergedJSON, err := json.Marshal(map1)
	if err != nil {
		return repository.JSON{}, err
	}

	return mergedJSON, nil
}

func AppendNewBeatmapStats(json1, json2 repository.JSON) (repository.JSON, error) {
	// merge two JSONs that are map[time.Time]model.BeatmapStatsModel
	map1 := make(model.BeatmapStats)
	map2 := make(model.BeatmapStats)

	if err := json.Unmarshal(json1, &map1); err != nil {
		return repository.JSON{}, err
	}

	if err := json.Unmarshal(json2, &map2); err != nil {
		return repository.JSON{}, err
	}

	for key, value := range map2 {
		map1[key] = value
	}

	mergedJSON, err := json.Marshal(map1)
	if err != nil {
		return repository.JSON{}, err
	}

	return mergedJSON, nil
}
