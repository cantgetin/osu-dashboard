package clean

import (
	"encoding/json"
	"osu-dashboard/internal/database/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRemoveAllMapEntriesExceptLastN_Mapset(t *testing.T) {
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
	in := json.RawMessage(jsonStats)

	n := 1
	expected := model.MapsetStats{
		sampleTime2: &model.MapsetStatsModel{
			Playcount: 2,
			Favorites: 2,
			Comments:  2,
		},
	}

	res, err := removeAllMapEntriesExceptLastN(&in, n)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	resMap := make(model.MapsetStats)
	err = json.Unmarshal(*res, &resMap)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	assert.Equal(t, expected, resMap)
}

func TestRemoveAllMapEntriesExceptLastN_Beatmap(t *testing.T) {
	sampleTime1 := time.Date(2023, 12, 23, 10, 0, 0, 0, time.UTC)
	sampleTime2 := time.Date(2023, 12, 24, 12, 0, 0, 0, time.UTC)

	stats := model.BeatmapStats{
		sampleTime1: &model.BeatmapStatsModel{
			Playcount: 1,
			Passcount: 3,
		},
		sampleTime2: &model.BeatmapStatsModel{
			Playcount: 1,
			Passcount: 4,
		},
	}

	jsonStats, err := json.Marshal(stats)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	in := json.RawMessage(jsonStats)

	n := 1
	expected := model.BeatmapStats{
		sampleTime2: &model.BeatmapStatsModel{
			Playcount: 1,
			Passcount: 4,
		},
	}

	res, err := removeAllMapEntriesExceptLastN(&in, n)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	resMap := make(model.BeatmapStats)
	err = json.Unmarshal(*res, &resMap)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	assert.Equal(t, expected, resMap)
}

func TestRemoveAllMapEntriesExceptLastN_User(t *testing.T) {
	sampleTime1 := time.Date(2023, 12, 23, 10, 0, 0, 0, time.UTC)
	sampleTime2 := time.Date(2023, 12, 24, 12, 0, 0, 0, time.UTC)

	stats := model.UserStats{
		sampleTime1: &model.UserStatsModel{
			PlayCount: 1,
			Favorites: 2,
			MapCount:  3,
			Comments:  4,
		},
		sampleTime2: &model.UserStatsModel{
			PlayCount: 2,
			Favorites: 3,
			MapCount:  4,
			Comments:  5,
		},
	}

	jsonStats, err := json.Marshal(stats)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	in := json.RawMessage(jsonStats)

	n := 1
	expected := model.UserStats{
		sampleTime2: &model.UserStatsModel{
			PlayCount: 2,
			Favorites: 3,
			MapCount:  4,
			Comments:  5,
		},
	}

	res, err := removeAllMapEntriesExceptLastN(&in, n)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	resMap := make(model.UserStats)
	err = json.Unmarshal(*res, &resMap)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	assert.Equal(t, expected, resMap)
}
