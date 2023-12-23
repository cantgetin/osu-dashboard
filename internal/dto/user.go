package dto

type CreateUserCommand struct {
	ID                       int    `json:"id"`
	AvatarURL                string `json:"avatar_url"`
	Username                 string `json:"username"`
	UnrankedBeatmapsetCount  int    `json:"unranked_beatmapset_count"`
	GraveyardBeatmapsetCount int    `json:"graveyard_beatmapset_count"`
}

type User struct {
	ID                       int    `json:"id"`
	AvatarURL                string `json:"avatar_url"`
	Username                 string `json:"username"`
	UnrankedBeatmapsetCount  int    `json:"unranked_beatmapset_count"`
	GraveyardBeatmapsetCount int    `json:"graveyard_beatmapset_count"`
}
