package dto

import (
	"playcount-monitor-backend/internal/database/repository/model"
	"time"
)

type CreateBeatmapCommand struct {
	Id               int       `json:"id"`
	BeatmapsetId     int       `json:"beatmapset_id"`
	DifficultyRating float64   `json:"difficulty_rating"`
	Version          string    `json:"version"`
	Accuracy         float64   `json:"accuracy"`
	Ar               float64   `json:"ar"`
	Bpm              float64   `json:"bpm"`
	Cs               float64   `json:"cs"`
	Status           string    `json:"status"`
	Url              string    `json:"url"`
	TotalLength      int       `json:"total_length"`
	UserId           int       `json:"user_id"`
	Passcount        int       `json:"passcount"`
	Playcount        int       `json:"playcount"`
	LastUpdated      time.Time `json:"last_updated"`
}

type Beatmap struct {
	Id               int                `json:"id"`
	BeatmapsetId     int                `json:"beatmapset_id"`
	DifficultyRating float64            `json:"difficulty_rating"`
	Version          string             `json:"version"`
	Accuracy         float64            `json:"accuracy"`
	Ar               float64            `json:"ar"`
	Bpm              float64            `json:"bpm"`
	Cs               float64            `json:"cs"`
	Status           string             `json:"status"`
	Url              string             `json:"url"`
	TotalLength      int                `json:"total_length"`
	UserId           int                `json:"user_id"`
	BeatmapStats     model.BeatmapStats `json:"beatmap_stats"`
	LastUpdated      time.Time          `json:"last_updated"`
}
