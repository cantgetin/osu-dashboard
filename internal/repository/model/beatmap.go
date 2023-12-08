package model

import (
	"playcount-monitor-backend/internal/repository"
	"time"
)

type Mapset struct {
	ID          int
	Artist      string
	Title       string
	Created     string
	Covers      map[string]string
	Status      string
	LastUpdated string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Beatmap struct {
	ID               int
	MapsetID         int
	DifficultyRating int
	Version          string
	BeatmapStats     repository.JSON
}

type BeatMapStats map[time.Time]MapStats

type MapStats struct {
	Playcount int `json:"plays"`
	Likes     int `json:"likes"`
}
