package model

import (
	"time"
)

type Mapset struct {
	ID          int
	Artist      string
	Title       string
	Created     string
	Covers      map[string]string
	Status      string
	LastUpdated string
	UserID      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
