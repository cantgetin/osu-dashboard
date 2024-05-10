package model

import "time"

type Following struct {
	ID          int
	Username    string
	CreatedAt   time.Time
	LastFetched time.Time
}
