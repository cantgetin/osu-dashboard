package mappers

import (
	"github.com/stretchr/testify/assert"
	"playcount-monitor-backend/internal/database/repository"
	"playcount-monitor-backend/internal/database/repository/model"
	"testing"
	"time"
)

// TODO: refactor
func Test_AppendNewMapsetStats(t *testing.T) {
	// Create sample JSON data
	json1 := repository.JSON(`{"2023-12-23T10:00:00Z":{"play_count":1,"favourite_count":1}}`)

	json2 := repository.JSON(`{"2023-12-24T12:00:00Z":{"play_count":2,"favourite_count":2}}`)

	expectedMergedJSON := repository.JSON(`{"2023-12-23T10:00:00Z":{"play_count":1,"favourite_count":1},"2023-12-24T12:00:00Z":{"play_count":2,"favourite_count":2}}`)

	mergedJSON, err := AppendNewMapsetStats(json1, json2)

	assert.NoError(t, err)
	assert.Equal(t, expectedMergedJSON, mergedJSON)
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
