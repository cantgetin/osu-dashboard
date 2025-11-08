package mappers

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"osu-dashboard/internal/database/repository"
	"osu-dashboard/internal/database/repository/model"
	"testing"
	"time"
)

// TODO: refactor
func Test_AppendNewMapsetStats(t *testing.T) {
	sampleTime1 := time.Date(2023, 12, 23, 10, 0, 0, 0, time.UTC)
	sampleTime2 := time.Date(2023, 12, 24, 12, 0, 0, 0, time.UTC)

	stats1 := model.MapsetStats{
		sampleTime1: &model.MapsetStatsModel{
			Playcount: 1,
			Favorites: 1,
			Comments:  1,
		},
	}

	stats2 := model.MapsetStats{
		sampleTime2: &model.MapsetStatsModel{
			Playcount: 2,
			Favorites: 2,
			Comments:  2,
		},
	}

	json1, err := json.Marshal(stats1)
	if err != nil {
		t.Error(err)
	}

	json2, err := json.Marshal(stats2)
	if err != nil {
		t.Error(err)
	}

	expectedMergedStats := model.MapsetStats{
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

	expectedMergedJSONBytes, err := json.Marshal(expectedMergedStats)
	if err != nil {
		t.Error(err)
	}

	expectedMergedJSON := repository.JSON(expectedMergedJSONBytes)

	mergedJSON, err := AppendNewMapsetStats(json1, json2)
	if err != nil {
		t.Error(err)
	}

	// Convert mergedJSON to repository.JSON
	actualMergedJSON := repository.JSON(mergedJSON)

	assert.NoError(t, err)
	assert.Equal(t, expectedMergedJSON, actualMergedJSON)
}

func Test_KeepLastNKeyValuesFromStats(t *testing.T) {
	data := model.UserStats{}
	for i := 0; i < 10; i++ {
		data[time.Now().UTC().AddDate(0, 0, -i).Truncate(time.Hour)] = &model.UserStatsModel{
			PlayCount: i,
			Favorites: i,
			MapCount:  i,
		}
	}

	KeepLastNKeyValuesFromStats(data, 7)
	assert.Equal(t, 7, len(data))
}
