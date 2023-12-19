package mapsetcreate

import "time"

type Covers struct {
	Cover       string `json:"cover"`
	Cover2X     string `json:"cover@2x"`
	Card        string `json:"card"`
	Card2X      string `json:"card@2x"`
	List        string `json:"list"`
	List2X      string `json:"list@2x"`
	Slimcover   string `json:"slimcover"`
	Slimcover2X string `json:"slimcover@2x"`
}

type Beatmap struct {
	BeatmapsetId     int         `json:"beatmapset_id"`
	DifficultyRating float64     `json:"difficulty_rating"`
	Id               int         `json:"id"`
	Mode             string      `json:"mode"`
	Status           string      `json:"status"`
	TotalLength      int         `json:"total_length"`
	UserId           int         `json:"user_id"`
	Version          string      `json:"version"`
	Accuracy         float64     `json:"accuracy"`
	Ar               float64     `json:"ar"`
	Bpm              int         `json:"bpm"`
	Convert          bool        `json:"convert"`
	CountCircles     int         `json:"count_circles"`
	CountSliders     int         `json:"count_sliders"`
	CountSpinners    int         `json:"count_spinners"`
	Cs               float64     `json:"cs"`
	DeletedAt        interface{} `json:"deleted_at"`
	Drain            float64     `json:"drain"`
	HitLength        int         `json:"hit_length"`
	IsScoreable      bool        `json:"is_scoreable"`
	LastUpdated      time.Time   `json:"last_updated"`
	ModeInt          int         `json:"mode_int"`
	Passcount        int         `json:"passcount"`
	Playcount        int         `json:"playcount"`
	Ranked           int         `json:"ranked"`
	Url              string      `json:"url"`
	Checksum         string      `json:"checksum"`
}

type CreateMapsetCommand struct {
	Key                string      `json:"key"`
	Artist             string      `json:"artist"`
	ArtistUnicode      string      `json:"artist_unicode"`
	Covers             Covers      `json:"covers"`
	Creator            string      `json:"creator"`
	FavouriteCount     int         `json:"favourite_count"`
	Hype               interface{} `json:"hype"`
	Id                 int         `json:"id"`
	Nsfw               bool        `json:"nsfw"`
	Offset             int         `json:"offset"`
	PlayCount          int         `json:"play_count"`
	PreviewUrl         string      `json:"preview_url"`
	Source             string      `json:"source"`
	Spotlight          bool        `json:"spotlight"`
	Status             string      `json:"status"`
	Title              string      `json:"title"`
	TitleUnicode       string      `json:"title_unicode"`
	TrackId            interface{} `json:"track_id"`
	UserId             int         `json:"user_id"`
	Video              bool        `json:"video"`
	Bpm                int         `json:"bpm"`
	CanBeHyped         bool        `json:"can_be_hyped"`
	DeletedAt          interface{} `json:"deleted_at"`
	DiscussionEnabled  bool        `json:"discussion_enabled"`
	DiscussionLocked   bool        `json:"discussion_locked"`
	IsScoreable        bool        `json:"is_scoreable"`
	LastUpdated        time.Time   `json:"last_updated"`
	LegacyThreadUrl    string      `json:"legacy_thread_url"`
	Ranked             int         `json:"ranked"`
	RankedDate         interface{} `json:"ranked_date"`
	Storyboard         bool        `json:"storyboard"`
	SubmittedDate      time.Time   `json:"submitted_date"`
	Tags               string      `json:"tags"`
	Beatmaps           []Beatmap   `json:"beatmaps"`
	NominationsSummary struct {
		Current  int `json:"current"`
		Required int `json:"required"`
	} `json:"nominations_summary"`
	Availability struct {
		DownloadDisabled bool        `json:"download_disabled"`
		MoreInformation  interface{} `json:"more_information"`
	} `json:"availability"`
}
