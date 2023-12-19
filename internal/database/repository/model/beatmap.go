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
	BeatmapStats     repository.JSON //BeatmapStats struct marshaled as JSON
}

type BeatmapStats map[time.Time]*BeatmapStatsModel

type BeatmapStatsModel struct {
	Playcount int `json:"plays"`
	Passcount int `json:"passes"`
}
