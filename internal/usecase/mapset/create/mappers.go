package mapsetcreate

import (
	json2 "encoding/json"
	"fmt"
	"playcount-monitor-backend/internal/database/repository"
	"playcount-monitor-backend/internal/database/repository/model"
	"time"
)

func mapBeatmapToBeatmapModel(beatmap *Beatmap) (*model.Beatmap, error) {
	stats, err := mapBeatmapStats(beatmap)
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

func mapBeatmapStats(beatmap *Beatmap) (repository.JSON, error) {
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

func mapCommandToMapsetModel(cmd *CreateMapsetCommand) (*model.Mapset, error) {
	mapsetStats, err := mapMapsetStats(cmd)
	if err != nil {
		return nil, err
	}

	return &model.Mapset{
		ID:          cmd.Id,
		Artist:      cmd.Artist,
		Title:       cmd.Title,
		Covers:      mapCovers(&cmd.Covers),
		Status:      cmd.Status,
		LastUpdated: cmd.LastUpdated,
		UserID:      cmd.UserId,
		Creator:     cmd.Creator,
		PreviewURL:  cmd.PreviewUrl,
		Tags:        cmd.Tags,
		MapsetStats: mapsetStats,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}, nil
}

func mapCovers(covers *Covers) map[string]string {
	return map[string]string{
		"cover":        covers.Cover,
		"cover@2x":     covers.Cover2X,
		"card":         covers.Card,
		"card@2x":      covers.Card2X,
		"list":         covers.List,
		"list@2x":      covers.List2X,
		"slimcover":    covers.Slimcover,
		"slimcover@2x": covers.Slimcover2X,
	}
}

func mapMapsetStats(cmd *CreateMapsetCommand) (repository.JSON, error) {
	var stats = make(model.MapsetStats)
	stats[time.Now().UTC()] = &model.MapsetStatsModel{
		Playcount: cmd.PlayCount,
		Favorites: cmd.FavouriteCount,
	}

	statsJson, err := json2.Marshal(stats)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal mapset stats: %w", err)
	}

	return statsJson, nil
}
