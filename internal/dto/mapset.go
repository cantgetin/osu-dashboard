package dto

import (
	"playcount-monitor-backend/internal/database/repository/model"
	"time"
)

type CreateMapsetCommand struct {
	Id             int                    `json:"id"`
	Artist         string                 `json:"artist"`
	Title          string                 `json:"title"`
	Covers         map[string]string      `json:"covers"`
	Status         string                 `json:"status"`
	LastUpdated    time.Time              `json:"last_updated"`
	UserId         int                    `json:"user_id"`
	PreviewUrl     string                 `json:"preview_url"`
	Tags           string                 `json:"tags"`
	PlayCount      int                    `json:"play_count"`
	FavouriteCount int                    `json:"favourite_count"`
	Bpm            int                    `json:"bpm"`
	Creator        string                 `json:"creator"`
	Beatmaps       []CreateBeatmapCommand `json:"beatmaps"`
}

type Mapset struct {
	Id          int               `json:"id"`
	Artist      string            `json:"artist"`
	Title       string            `json:"title"`
	Covers      map[string]string `json:"covers"`
	Status      string            `json:"status"`
	LastUpdated time.Time         `json:"last_updated"`
	UserId      int               `json:"user_id"`
	PreviewUrl  string            `json:"preview_url"`
	Tags        string            `json:"tags"`
	MapsetStats model.MapsetStats `json:"mapset_stats"`
	Bpm         int               `json:"bpm"`
	Creator     string            `json:"creator"`
	Beatmaps    []Beatmap         `json:"beatmaps"`
}
