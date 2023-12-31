package dto

import "time"

type Tracking struct {
	ID            int       `json:"id"`
	Username      string    `json:"username"`
	TrackingSince time.Time `json:"tracking_since"`
}
