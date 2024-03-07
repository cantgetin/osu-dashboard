package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"playcount-monitor-backend/internal/database/repository/model"
	"testing"
	"time"
)

// todo: figure this out
func TestRemoveAllMapEntriesExceptLastN(t *testing.T) {
	t.Skip("TODO: fix")

	sampleTime1 := time.Date(2023, 12, 23, 10, 0, 0, 0, time.UTC)
	sampleTime2 := time.Date(2023, 12, 24, 12, 0, 0, 0, time.UTC)

	stats := model.MapsetStats{
		sampleTime1: &model.MapsetStatsModel{
			Playcount: 1,
			Favorites: 1,
			Comments:  1,
		},
		sampleTime2: &model.MapsetStatsModel{
			Playcount: 2,
			Favorites: 2,
			Comments:  2,
		},
	}

	jsonStats, err := json.Marshal(stats)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	n := 1
	expected := model.MapsetStats{
		sampleTime2: &model.MapsetStatsModel{
			Playcount: 2,
			Favorites: 2,
			Comments:  2,
		},
	}

	err = removeAllMapEntriesExceptLastN(jsonStats, n)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	assert.Equal(t, expected, jsonStats)
}
