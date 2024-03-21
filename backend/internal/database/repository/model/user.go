package model

import (
	"playcount-monitor-backend/internal/database/repository"
	"time"
)

type User struct {
	ID                       int
	Username                 string
	AvatarURL                string
	GraveyardBeatmapsetCount int
	UnrankedBeatmapsetCount  int
	UserStats                repository.JSON
	MapCounts                repository.JSON
	CreatedAt                time.Time
	UpdatedAt                time.Time
}

type UserStats map[time.Time]*UserStatsModel

type UserStatsModel struct {
	PlayCount int `json:"play_count"`
	Favorites int `json:"favourite_count"`
	MapCount  int `json:"map_count"`
	Comments  int `json:"comments_count"`
}

type MapCounts struct {
	Graveyard int `json:"graveyard"`
	WIP       int `json:"wip"`
	Pending   int `json:"pending"`
	Ranked    int `json:"ranked"`
	Approved  int `json:"approved"`
	Qualified int `json:"qualified"`
	Loved     int `json:"loved"`
}
