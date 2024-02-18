package dto

import (
	"playcount-monitor-backend/internal/database/repository/model"
	"time"
)

type User struct {
	ID            int             `json:"id"`
	AvatarURL     string          `json:"avatar_url"`
	Username      string          `json:"username"`
	TrackingSince time.Time       `json:"tracking_since"`
	UserStats     model.UserStats `json:"user_stats"`
	UserMapCounts *UserMapCounts  `json:"user_map_counts"`
}

type UserMapCounts struct {
	Graveyard int `json:"graveyard"`
	Wip       int `json:"wip"`
	Pending   int `json:"pending"`
	Ranked    int `json:"ranked"`
	Approved  int `json:"approved"`
	Qualified int `json:"qualified"`
	Loved     int `json:"loved"`
}
