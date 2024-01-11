package osuapi

import "time"

type User struct {
	ID                       int    `json:"id"`
	AvatarURL                string `json:"avatar_url"`
	Username                 string `json:"username"`
	UnrankedBeatmapsetCount  int    `json:"unranked_beatmapset_count"`
	GraveyardBeatmapsetCount int    `json:"graveyard_beatmapset_count"`
}

type Mapset struct {
	Id             int               `json:"id"`
	Artist         string            `json:"artist"`
	Title          string            `json:"title"`
	Covers         map[string]string `json:"covers"`
	Status         string            `json:"status"`
	LastUpdated    time.Time         `json:"last_updated"`
	UserId         int               `json:"user_id"`
	PreviewUrl     string            `json:"preview_url"`
	Tags           string            `json:"tags"`
	PlayCount      int               `json:"play_count"`
	FavouriteCount int               `json:"favourite_count"`
	Bpm            float64           `json:"bpm"`
	Creator        string            `json:"creator"`
	Beatmaps       []*Beatmap        `json:"beatmaps"`
}

type Beatmap struct {
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
