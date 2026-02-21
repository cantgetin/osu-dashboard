package model

import "time"

type Job struct {
	ID        int
	Name      string
	StartedAt time.Time
	EndedAt   time.Time
	Error     bool
	ErrorText string
}
