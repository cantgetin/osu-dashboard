package mappers

import (
	"encoding/json"
	"osu-dashboard/internal/database/repository"
	"osu-dashboard/internal/database/repository/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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
	assert.Len(t, data, 7)
}

func TestGetMapsetCover(t *testing.T) {
	type args struct {
		mapset   *model.Mapset
		coverKey string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				mapset: &model.Mapset{
					Covers: []byte(`{"card":"https://assets.ppy.sh/beatmaps/2509233/covers/list.jpg?1771254436"}`),
				},
				coverKey: "card",
			},
			want: "https://assets.ppy.sh/beatmaps/2509233/covers/list.jpg?1771254436",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, GetMapsetCover(tt.args.mapset, tt.args.coverKey),
				"GetMapsetCover(%v, %v)", tt.args.mapset, tt.args.coverKey)
		})
	}
}
