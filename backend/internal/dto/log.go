package dto

import "time"

type Log struct {
	ID                 int           `json:"id,omitempty"`
	Name               string        `json:"name,omitempty"`
	Message            string        `json:"message,omitempty"`
	Service            string        `json:"service,omitempty"`
	AppVersion         string        `json:"app_version,omitempty"`
	Platform           string        `json:"platform,omitempty"`
	Type               string        `json:"type,omitempty"`
	APIRequests        int           `json:"api_requests,omitempty"`
	SuccessRatePercent float64       `json:"success_rate_percent,omitempty"`
	TrackedAt          time.Time     `json:"tracked_at"`
	AvgResponseTime    time.Duration `json:"avg_response_time,omitempty"`
	ElapsedTime        time.Duration `json:"elapsed_time,omitempty"`
	TimeSinceLastTrack time.Duration `json:"time_since_last_track,omitempty"`
}

type LogsPaged struct {
	Logs        []*Log `json:"logs,omitempty"`
	CurrentPage int    `json:"current_page,omitempty"`
	Pages       int    `json:"pages,omitempty"`
}
