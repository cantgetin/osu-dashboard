package osuapimodels

import "time"

type Mapset struct {
	Id            int               `json:"id"`
	Artist        string            `json:"artist"`
	Title         string            `json:"title"`
	Covers        map[string]string `json:"covers"`
	Status        string            `json:"status"`
	LastUpdated   time.Time         `json:"last_updated"`
	UserId        int               `json:"user_id"`
	PreviewUrl    string            `json:"preview_url"`
	Tags          string            `json:"tags"`
	PlayCount     int               `json:"play_count"`
	FavoriteCount int               `json:"favourite_count"` // do not rename this
	Bpm           float64           `json:"bpm"`
	Creator       string            `json:"creator"`
	Beatmaps      []*Beatmap        `json:"beatmaps"`
}

type MapsetExtended struct {
	CommentsCount int    `json:"comments_count"`
	Genre         string `json:"genre"`
	Language      string `json:"language"`
	*Mapset
}

type MapsetLangGenre struct {
	Genre struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"genre"`
	Language struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"language"`
}
