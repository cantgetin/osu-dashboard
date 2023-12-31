package model

import "time"

type Tracking struct {
	ID        int
	Username  string
	CreatedAt time.Time
}
