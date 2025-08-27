package dto

import "time"

type Log struct {
	ID                 int           `json:"id"`
	Name               string        `json:"name"`
	Message            string        `json:"message"`
	Service            string        `json:"service"`
	AppVersion         string        `json:"app_version"`
	Platform           string        `json:"platform"`
	Type               string        `json:"type"`
	APIRequests        int           `json:"api_requests"`
	SuccessRatePercent float64       `json:"success_rate_percent"`
	TrackedAt          time.Time     `json:"tracked_at"`
	AvgResponseTime    time.Duration `json:"avg_response_time"`
	ElapsedTime        time.Duration `json:"elapsed_time"`
	TimeSinceLastTrack time.Duration `json:"time_since_last_track"`
}

type LogsPaged struct {
	Logs        []Log `json:"logs"`
	CurrentPage int   `json:"current_page"`
	Pages       int   `json:"pages"`
}
