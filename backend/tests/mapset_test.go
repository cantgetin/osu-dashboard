package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"playcount-monitor-backend/internal/app/mapsetserviceapi"
	"playcount-monitor-backend/internal/database/repository"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/dto"
	"playcount-monitor-backend/tests/integration"
	"time"
)

func (s *IntegrationSuite) Test_ListMapsets() {
	s.Run("valid requests", func() {
		type models struct {
			User     *model.User
			Mapsets  []*model.Mapset
			Beatmaps []*model.Beatmap
		}

		var tt = []struct {
			name    string
			create  *models
			out     *mapsetserviceapi.MapsetListResponse
			outCode int
		}{
			{
				name: "1 user, 2 mapsets with 1 beatmap each",
				create: &models{
					User: &model.User{
						ID:                       1,
						AvatarURL:                "avatarurl.com",
						Username:                 "username",
						UnrankedBeatmapsetCount:  1,
						GraveyardBeatmapsetCount: 1,
						UserStats:                repository.JSON(`{"2023-12-24T12:00:00Z":{"play_count":100,"favourite_count":2, "map_count":1}}`),
					},
					Mapsets: []*model.Mapset{
						{
							ID:          1,
							Artist:      "artist",
							Title:       "title",
							Covers:      repository.JSON(`{"cover1": "cover1", "cover2": "cover2"}`),
							Status:      "graveyard",
							LastUpdated: time.Now().UTC(),
							UserID:      1,
							Creator:     "username",
							PreviewURL:  "previewurl.com",
							Tags:        "tags shmags",
							BPM:         210,
							MapsetStats: repository.JSON(`{"2023-12-24T12:00:00Z":{"play_count":100,"favourite_count":2}}`),
						},
						{
							ID:          2,
							Artist:      "artist2",
							Title:       "title2",
							Covers:      repository.JSON(`{"cover1": "cover1", "cover2": "cover2"}`),
							Status:      "graveyard",
							LastUpdated: time.Now().UTC(),
							UserID:      1,
							Creator:     "username",
							PreviewURL:  "previewurl.com",
							Tags:        "tags shmags",
							BPM:         220,
							MapsetStats: repository.JSON(`{"2023-12-24T12:00:00Z":{"play_count":100,"favourite_count":2}}`),
						},
					},
					Beatmaps: []*model.Beatmap{
						{
							ID:               1,
							MapsetID:         1,
							DifficultyRating: 5.3,
							Version:          "version1",
							Accuracy:         3,
							AR:               8.5,
							BPM:              210,
							CS:               4.3,
							Status:           "graveyard",
							URL:              "url.com",
							TotalLength:      100,
							UserID:           1,
							BeatmapStats:     repository.JSON(`{"2023-12-24T12:00:00Z":{"play_count":52,"pass_count":2}}`),
						},
						{
							ID:               2,
							MapsetID:         2,
							DifficultyRating: 5.4,
							Version:          "version2",
							Accuracy:         5,
							AR:               7.5,
							BPM:              210,
							CS:               3.3,
							Status:           "graveyard",
							URL:              "url2.com",
							TotalLength:      102,
							UserID:           1,
							BeatmapStats:     repository.JSON(`{"2023-12-24T12:00:00Z":{"play_count":13,"pass_count":2}}`),
						},
					},
				},
				out: &mapsetserviceapi.MapsetListResponse{
					Mapsets: []*dto.Mapset{
						{
							Id:          2,
							Artist:      "artist2",
							Title:       "title2",
							Covers:      map[string]string{"cover1": "cover1", "cover2": "cover2"},
							Status:      "graveyard",
							LastUpdated: time.Now().UTC(),
							UserId:      1,
							PreviewUrl:  "previewurl.com",
							Tags:        "tags shmags",
							Bpm:         220,
							Creator:     "username",
							Beatmaps: []*dto.Beatmap{
								{
									Id:               2,
									BeatmapsetId:     2,
									DifficultyRating: 5.4,
									Version:          "version2",
									Accuracy:         5,
									Ar:               7.5,
									Bpm:              210,
									Cs:               3.3,
									Status:           "graveyard",
									Url:              "url2.com",
									TotalLength:      102,
									UserId:           1,
								},
							},
						},
						{
							Id:          1,
							Artist:      "artist",
							Title:       "title",
							Covers:      map[string]string{"cover1": "cover1", "cover2": "cover2"},
							Status:      "graveyard",
							LastUpdated: time.Now().UTC(),
							UserId:      1,
							PreviewUrl:  "previewurl.com",
							Tags:        "tags shmags",
							Bpm:         210,
							Creator:     "username",
							Beatmaps: []*dto.Beatmap{
								{
									Id:               1,
									BeatmapsetId:     1,
									DifficultyRating: 5.3,
									Version:          "version1",
									Accuracy:         3,
									Ar:               8.5,
									Bpm:              210,
									Cs:               4.3,
									Status:           "graveyard",
									Url:              "url.com",
									TotalLength:      100,
									UserId:           1,
								},
								{
									Id:               2,
									BeatmapsetId:     1,
									DifficultyRating: 5.4,
									Version:          "version2",
									Accuracy:         5,
									Ar:               7.5,
									Bpm:              210,
									Cs:               3.3,
									Status:           "graveyard",
									Url:              "url2.com",
									TotalLength:      102,
									UserId:           1,
								},
							},
						},
					},
					CurrentPage: 1,
					Pages:       1,
				},
				outCode: 200,
			},
		}

		for _, tc := range tt {
			s.Run(tc.name, func() {
				// create models
				err := s.db.Create(&tc.create.User).Error
				s.Require().NoError(err)

				for _, ms := range tc.create.Mapsets {
					err := s.db.Create(&ms).Error
					s.Require().NoError(err)
				}

				for _, bm := range tc.create.Beatmaps {
					err := s.db.Create(&bm).Error
					s.Require().NoError(err)
				}

				url := fmt.Sprintf("http://localhost:%s/api/beatmapset/list?sort=created_at&order=desc", s.port)
				out, err := http.Get(url)

				s.Require().NoError(err)
				s.Require().Equal(tc.outCode, out.StatusCode)

				defer out.Body.Close()

				body, err := io.ReadAll(out.Body)
				s.Require().NoError(err)

				var actual mapsetserviceapi.MapsetListResponse
				err = json.Unmarshal(body, &actual)
				s.Require().NoError(err)

				s.Assert().Equal(2, len(actual.Mapsets))

				for i, actualMapset := range actual.Mapsets {
					expectedMapset := tc.out.Mapsets[i]

					s.Assert().Equal(expectedMapset.Id, actualMapset.Id)
					s.Assert().Equal(expectedMapset.Artist, actualMapset.Artist)
					s.Assert().Equal(expectedMapset.Title, actualMapset.Title)
					s.Assert().Equal(expectedMapset.Covers, actualMapset.Covers)
					s.Assert().Equal(expectedMapset.Status, actualMapset.Status)
					s.Assert().Equal(expectedMapset.UserId, actualMapset.UserId)
					s.Assert().Equal(expectedMapset.PreviewUrl, actualMapset.PreviewUrl)
					s.Assert().Equal(expectedMapset.Tags, actualMapset.Tags)
					s.Assert().Equal(expectedMapset.Bpm, actualMapset.Bpm)
					s.Assert().Equal(expectedMapset.Creator, actualMapset.Creator)

					s.Assert().Equal(1, len(actualMapset.MapsetStats))

					for j, actualBeatmap := range actualMapset.Beatmaps {
						expectedBeatmap := tc.out.Mapsets[i].Beatmaps[j]

						s.Assert().Equal(expectedBeatmap.Id, actualBeatmap.Id)
						s.Assert().Equal(expectedBeatmap.BeatmapsetId, actualBeatmap.BeatmapsetId)
						s.Assert().Equal(expectedBeatmap.DifficultyRating, actualBeatmap.DifficultyRating)
						s.Assert().Equal(expectedBeatmap.Version, actualBeatmap.Version)
						s.Assert().Equal(expectedBeatmap.Accuracy, actualBeatmap.Accuracy)
						s.Assert().Equal(expectedBeatmap.Ar, actualBeatmap.Ar)
						s.Assert().Equal(expectedBeatmap.Bpm, actualBeatmap.Bpm)
						s.Assert().Equal(expectedBeatmap.Cs, actualBeatmap.Cs)
						s.Assert().Equal(expectedBeatmap.Status, actualBeatmap.Status)
						s.Assert().Equal(expectedBeatmap.Url, actualBeatmap.Url)
						s.Assert().Equal(expectedBeatmap.TotalLength, actualBeatmap.TotalLength)
						s.Assert().Equal(expectedBeatmap.UserId, actualBeatmap.UserId)

						s.Assert().Equal(1, len(actualBeatmap.BeatmapStats))
					}
				}
			})
		}

		err := integration.ClearTables(s.ctx, s.db)
		if err != nil {
			return
		}
	})
	err := integration.ClearTables(s.ctx, s.db)
	if err != nil {
		s.T().Fatal(err)
	}
}
