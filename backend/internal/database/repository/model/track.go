package model

import "time"

type Track struct {
	ID        int `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	TrackedAt time.Time
}
