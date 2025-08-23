package model

import "time"

type LogType string

const (
	TrackTypeInitial LogType = "single"
	TrackTypeRegular LogType = "regular"
)

type Log struct {
	ID                 int           `gorm:"primaryKey;autoIncrement"`
	Name               string        `gorm:"type:text;not null"`
	Message            LogMessage    `gorm:"type:text;not null"`
	Service            string        `gorm:"type:text;not null"`
	AppVersion         string        `gorm:"type:text;not null"`
	Platform           string        `gorm:"type:text;not null"`
	Type               LogType       `gorm:"type:text;not null"`
	APIRequests        int           `gorm:"not null"`
	SuccessRatePercent float64       `gorm:"not null"`
	TrackedAt          time.Time     `gorm:"type:timestamp;not null"`
	AvgResponseTime    time.Duration `gorm:"not null"`
	ElapsedTime        time.Duration `gorm:"not null"`
	TimeSinceLastTrack time.Duration `gorm:"not null"`
}

type LogMessage string

const (
	LogMessageDailyTrack   LogMessage = "Adding new statistic for all added users: tracking playcount, comment, favourite increases, new mapsets, beatmaps, tags, avatars etc."
	LogMessageDailyClean   LogMessage = "Cleaning old statistic records for all added users: clean stats records of mapsets, beatmaps, users, only keep records for last two weeks."
	LogMessageInitialTrack LogMessage = "Adding initial statistic for user that joined: tracking plays, comments, favourites, mapsets, beatmaps, tags, avatar etc."
)
