package model

import "time"

type User struct {
	ID                       int
	Username                 string
	AvatarURL                string
	GraveyardBeatmapsetCount int
	UnrankedBeatmapsetCount  int
	CreatedAt                time.Time
	UpdatedAt                time.Time
}
