package dto

import "time"

type Beatmap struct {
	Id               int       `json:"id"`
	BeatmapsetId     int       `json:"beatmapset_id"`
	DifficultyRating float64   `json:"difficulty_rating"`
	Version          string    `json:"version"`
	Accuracy         float64   `json:"accuracy"`
	Ar               float64   `json:"ar"`
	Bpm              int       `json:"bpm"`
	Cs               float64   `json:"cs"`
	Status           string    `json:"status"`
	Url              string    `json:"url"`
	TotalLength      int       `json:"total_length"`
	UserId           int       `json:"user_id"`
	Passcount        int       `json:"passcount"`
	Playcount        int       `json:"playcount"`
	LastUpdated      time.Time `json:"last_updated"`

	// unused
	// Mode string `json:"mode"`
	// Convert       bool `json:"convert"`
	// CountCircles  int  `json:"count_circles"`
	// CountSliders  int  `json:"count_sliders"`
	// CountSpinners int  `json:"count_spinners"`
	// DeletedAt   interface{} `json:"deleted_at"`
	// Drain       float64     `json:"drain"`
	// HitLength   int         `json:"hit_length"`
	// IsScoreable bool        `json:"is_scoreable"`
	// ModeInt int `json:"mode_int"`
	// Ranked int `json:"ranked"`
	// Checksum string `json:"checksum"`
}
