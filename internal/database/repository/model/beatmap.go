package model

import (
	"playcount-monitor-backend/internal/database/repository"
	"time"
)

type Beatmap struct {
	ID               int
	MapsetID         int
	DifficultyRating float64
	Version          string // diff name
	Accuracy         float64
	AR               float64
	BPM              float64
	CS               float64
	Status           string
	URL              string
	TotalLength      int
	UserID           int
	LastUpdated      time.Time       // last map update
	BeatmapStats     repository.JSON `gorm:"type:jsonb"` //BeatmapStats struct marshaled as JSON
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type BeatmapStats map[time.Time]*BeatmapStatsModel

type BeatmapStatsModel struct {
	Playcount int `json:"play_count"`
	Passcount int `json:"pass_count"`
}
