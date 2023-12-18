package model

import "time"

type User struct {
	ID                       int
	AvatarURL                string
	Username                 string
	GraveyardBeatmapsetCount int
	UnrankedBeatmapsetCount  int
	CreatedAt                time.Time
	UpdatedAt                time.Time
}
