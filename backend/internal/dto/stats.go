package dto

type UserMapStatistics struct {
	Tags      map[string]int `json:"most_popular_tags"`
	Languages map[string]int `json:"most_popular_languages"`
	Genres    map[string]int `json:"most_popular_genres"`
	BPMs      map[string]int `json:"most_popular_bpms"`
	Starrates map[string]int `json:"most_popular_starrates"`
}

type SystemStatistics struct {
	Users    int `json:"users"`
	Mapsets  int `json:"mapsets"`
	Beatmaps int `json:"beatmaps"`
	Tracks   int `json:"tracks"`
}
