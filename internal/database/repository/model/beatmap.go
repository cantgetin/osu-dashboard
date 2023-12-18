package model

import (
	"playcount-monitor-backend/internal/database/repository"
	"time"
)

type Beatmap struct {
	ID               int
	MapsetID         int
	DifficultyRating int
	Version          string
	BeatmapStats     repository.JSON //BeatmapStats struct marshaled as JSON
}

type BeatmapStats map[time.Time]*MapStats

type MapStats struct {
	Playcount int `json:"plays"`
	Likes     int `json:"likes"`
}
