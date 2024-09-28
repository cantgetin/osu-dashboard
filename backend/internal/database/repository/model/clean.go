package model

import "time"

type Clean struct {
	ID        int `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	CleanedAt time.Time
}
