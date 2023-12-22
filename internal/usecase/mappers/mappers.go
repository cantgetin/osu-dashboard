package mappers

import (
	json2 "encoding/json"
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

	statsJson, err := json2.Marshal(stats)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal beatmap stats: %w", err)
	}

	return statsJson, nil
}

func mapCovers(m map[string]string) (repository.JSON, error) {
	coversJson, err := json2.Marshal(m)
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

	statsJson, err := json2.Marshal(stats)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal mapset stats to json: %w", err)
	}

	return statsJson, nil
}
