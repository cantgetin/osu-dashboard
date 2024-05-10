package dto

import "time"

type Following struct {
	ID             int       `json:"id"`
	Username       string    `json:"username"`
	FollowingSince time.Time `json:"following_since"`
	LastFetched    time.Time `json:"last_fetched"`
}
