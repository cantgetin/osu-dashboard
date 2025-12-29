package dto

import (
	"osu-dashboard/internal/database/repository/model"
	"time"
)

type Mapset struct {
	Id          int               `json:"id"`
	Artist      string            `json:"artist"`
	Title       string            `json:"title"`
	Covers      map[string]string `json:"covers"`
	Status      string            `json:"status"`
	Genre       string            `json:"genre"`
	Language    string            `json:"language"`
	LastUpdated time.Time         `json:"last_updated"`
	UserId      int               `json:"user_id"`
	PreviewUrl  string            `json:"preview_url"`
	Tags        string            `json:"tags"`
	MapsetStats model.MapsetStats `json:"mapset_stats"`
	Bpm         float64           `json:"bpm"`
	Creator     string            `json:"creator"`
	Beatmaps    []*Beatmap        `json:"beatmaps"`
}

type MapsetsPaged struct {
	Mapsets     []*Mapset `json:"mapsets"`
	CurrentPage int       `json:"current_page"`
	Pages       int       `json:"pages"`
}
