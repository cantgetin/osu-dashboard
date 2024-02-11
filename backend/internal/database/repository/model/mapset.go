package model

import (
	"playcount-monitor-backend/internal/database/repository"
	"time"
)

type Mapset struct {
	ID            int
	Artist        string
	Title         string
	Covers        repository.JSON `gorm:"type:jsonb"` // map[string]string probably
	Status        string
	LastUpdated   time.Time
	UserID        int
	Creator       string
	PreviewURL    string
	Tags          string
	BPM           float64
	MapsetStats   repository.JSON `gorm:"type:jsonb"` // MapsetStats struct marshaled as JSON
	LastPlaycount int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type MapsetStats map[time.Time]*MapsetStatsModel

type MapsetStatsModel struct {
	Playcount int `json:"play_count"`
	Favorites int `json:"favourite_count"`
}
