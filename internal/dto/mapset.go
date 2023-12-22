package dto

import "time"

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
	Bpm            int               `json:"bpm"`
	Creator        string            `json:"creator"`
	Beatmaps       []Beatmap         `json:"beatmaps"`

	// unused
	// Key           string      `json:"key"`
	// ArtistUnicode string      `json:"artist_unicode"`
	// Hype          interface{} `json:"hype"`
	// Nsfw          bool        `json:"nsfw"`
	// Offset        int         `json:"offset"`
	// Source        string      `json:"source"`
	// Spotlight     bool        `json:"spotlight"`
	// TitleUnicode  string      `json:"title_unicode"`
	// TrackId       interface{} `json:"track_id"`
	// Video         bool        `json:"video"`
	// CanBeHyped         bool        `json:"can_be_hyped"`
	// DeletedAt          interface{} `json:"deleted_at"`
	// DiscussionEnabled  bool        `json:"discussion_enabled"`
	// DiscussionLocked   bool        `json:"discussion_locked"`
	// IsScoreable        bool        `json:"is_scoreable"`
	// LegacyThreadUrl    string      `json:"legacy_thread_url"`
	// Ranked             int         `json:"ranked"`
	// RankedDate         interface{} `json:"ranked_date"`
	// Storyboard         bool        `json:"storyboard"`
	// SubmittedDate      time.Time   `json:"submitted_date"`
	// Beatmaps           []Beatmap   `json:"beatmaps"`
	// NominationsSummary struct {
	// 	Current  int `json:"current"`
	// 	Required int `json:"required"`
	// } `json:"nominations_summary"`
	// Availability struct {
	// 	DownloadDisabled bool        `json:"download_disabled"`
	// 	MoreInformation  interface{} `json:"more_information"`
	// } `json:"availability"`
}
