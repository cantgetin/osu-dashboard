package mappers

import (
	"encoding/json"
	"fmt"
	"playcount-monitor-backend/internal/database/repository"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/dto"
	"time"
)

// command -> model

func MapCreateBeatmapCommandToBeatmapModel(beatmap *dto.CreateBeatmapCommand) (*model.Beatmap, error) {
	stats, err := MapCreateBeatmapCommandToStatsJSON(beatmap)
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
		BPM:              beatmap.Bpm,
		CS:               beatmap.Cs,
		Status:           beatmap.Status,
		URL:              beatmap.Url,
		TotalLength:      beatmap.TotalLength,
		UserID:           beatmap.UserId,
		LastUpdated:      beatmap.LastUpdated,
		BeatmapStats:     stats,
		CreatedAt:        time.Now().UTC(),
		UpdatedAt:        time.Now().UTC(),
	}, nil
}

func MapCreateUserCommandToUserModel(user *dto.CreateUserCommand) (*model.User, error) {
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

func MapCreateMapsetCommandToMapsetModel(mapset *dto.CreateMapsetCommand) (*model.Mapset, error) {
	mapsetStats, err := MapCreateMapsetCommandToStatsJSON(mapset)
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

// model -> dto

func MapUserModelToUserDTO(user *model.User) *dto.User {
	return &dto.User{
		ID:                       user.ID,
		Username:                 user.Username,
		AvatarURL:                user.AvatarURL,
		GraveyardBeatmapsetCount: user.GraveyardBeatmapsetCount,
		UnrankedBeatmapsetCount:  user.UnrankedBeatmapsetCount,
	}
}

func MapBeatmapModelToBeatmapDTO(beatmap *model.Beatmap) (*dto.Beatmap, error) {
	stats, err := MapStatsJSONToBeatmapStats(beatmap.BeatmapStats)
	if err != nil {
		return nil, err
	}

	return &dto.Beatmap{
		Id:               beatmap.ID,
		BeatmapsetId:     beatmap.MapsetID,
		DifficultyRating: beatmap.DifficultyRating,
		Version:          beatmap.Version,
		Accuracy:         beatmap.Accuracy,
		Ar:               beatmap.AR,
		Bpm:              beatmap.BPM,
		Cs:               beatmap.CS,
		Status:           beatmap.Status,
		Url:              beatmap.URL,
		TotalLength:      beatmap.TotalLength,
		UserId:           beatmap.UserID,
		LastUpdated:      beatmap.LastUpdated,
		BeatmapStats:     stats,
	}, nil
}

func MapMapsetModelToMapsetDTO(mapset *model.Mapset, beaatmaps []*model.Beatmap) *dto.Mapset {

	for _, beatmap := range beaatmaps {
		bm := MapBeatmapModelToBeatmapDTO(beatmap)
	}
	return &dto.Mapset{
		Id:             mapset.ID,
		Artist:         mapset.Artist,
		Title:          mapset.Title,
		Covers:         mapset.Covers,
		Status:         mapset.Status,
		LastUpdated:    mapset.LastUpdated,
		UserId:         mapset.UserID,
		PreviewUrl:     mapset.PreviewURL,
		Tags:           mapset.Tags,
		PlayCount:      0,
		FavouriteCount: 0,
		Bpm:            0,
		Creator:        mapset.Creator,
		Beatmaps:       nil,
	}
}

func MapCreateBeatmapCommandToStatsJSON(beatmap *dto.CreateBeatmapCommand) (repository.JSON, error) {
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

func MapCreateMapsetCommandToStatsJSON(mapset *dto.CreateMapsetCommand) (repository.JSON, error) {
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

func MapStatsJSONToMapsetStats(json repository.JSON) (*dto.Mapset, error) {
	var mapset *dto.Mapset

	var mapsetStats model.MapsetStats

	if err := json.Unmarshal(json, &mapsetStats); err != nil {
		return nil, fmt.Errorf("failed to unmarshal mapset stats: %w", err)
	}

	return &dto.Mapset{}
}

func MapStatsJSONToBeatmapStats(json repository.JSON) (model.BeatmapStats, error) {
	beatmapStats := make(model.BeatmapStats)
	if err := json.Unmarshal(json, &beatmapStats); err != nil {
		return nil, fmt.Errorf("failed to unmarshal beatmap stats: %w", err)
	}

	return beatmapStats, nil
}

func mapCovers(m map[string]string) (repository.JSON, error) {
	coversJson, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal mapset covers to json: %w", err)
	}

	return coversJson, nil
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
