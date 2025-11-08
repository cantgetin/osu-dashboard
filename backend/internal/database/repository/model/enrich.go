package model

import "time"

type Enrich struct {
	ID         int `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	EnrichedAt time.Time
}
