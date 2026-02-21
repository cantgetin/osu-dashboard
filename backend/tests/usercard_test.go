package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"osu-dashboard/internal/database/model"
	"osu-dashboard/internal/dto"
	"osu-dashboard/internal/usecase/command"
	"osu-dashboard/tests/integration"
	"time"
)

func (s *IntegrationSuite) Test_CreateUseCard() {
	s.T().Skip("this endpoint disabled for now")
	// todo: might test usecase later

	s.Run("valid requests", func() {
		tt := []struct {
			name    string
			in      *command.CreateUserCardCommand
			outCode int
		}{
			{
				name: "valid request, should properly create",
				in: &command.CreateUserCardCommand{
					User: &command.CreateUserCommand{
						ID:        1,
						AvatarURL: "avatarurl.com",
						Username:  "username",
					},
					Mapsets: []*command.CreateMapsetCommand{
						{
							Id:     1,
							Artist: "artist",
							Title:  "title",
							Covers: map[string]string{
								"cover1": "cover1",
								"cover2": "cover2",
							},
							Status:        "graveyard",
							LastUpdated:   time.Now().UTC(),
							UserId:        1,
							PreviewUrl:    "previewurl.com",
							Tags:          "tags tags",
							PlayCount:     20,
							FavoriteCount: 25,
							Bpm:           150,
							Creator:       "username",
							Beatmaps: []*command.CreateBeatmapCommand{
								{
									Id:               1,
									BeatmapsetId:     1,
									DifficultyRating: 5.3,
									Version:          "version1",
									Accuracy:         6.7,
									Ar:               9.3,
									Bpm:              150.3,
									Cs:               4,
									Status:           "graveyard",
									Url:              "beatmapurl.com",
									TotalLength:      3,
									UserId:           1,
									Passcount:        12,
									Playcount:        13,
									LastUpdated:      time.Now().UTC(),
								},
								{
									Id:               2,
									BeatmapsetId:     1,
									DifficultyRating: 6.8,
									Version:          "version2",
									Accuracy:         4.6,
									Ar:               9,
									Bpm:              150,
									Cs:               5,
									Status:           "graveyard",
									Url:              "beatmap2url.com",
									TotalLength:      4,
									UserId:           1,
									Passcount:        0,
									Playcount:        7,
									LastUpdated:      time.Now().UTC(),
								},
							},
						},
					},
				},
				outCode: 200,
			},
		}

		for _, tc := range tt {
			s.Run(tc.name, func() {
				inJSON, err := json.Marshal(tc.in)
				s.Require().NoError(err)

				out, err := http.Post(
					fmt.Sprintf("http://localhost:%s/api/user_card/create", s.port),
					"application/json",
					bytes.NewBuffer(inJSON),
				)
				defer out.Body.Close()

				s.Require().NoError(err)
				s.Require().NotNil(out)

				s.Equal(tc.outCode, out.StatusCode)
			})
		}
	})
	err := integration.ClearTables(s.ctx, s.db)
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *IntegrationSuite) Test_UpdateUserCard() {
	s.T().Skip("this endpoint disabled for now")
	// todo: might test usecase later

	s.Run("valid requests", func() {
		type models struct {
			User     *model.User
			Mapsets  []*model.Mapset
			Beatmaps []*model.Beatmap
		}

		var tt = []struct {
			name    string
			create  *models
			in      *command.UpdateUserCardCommand
			result  *models // assert db models after calling update method
			outCode int
		}{
			{
				name: "valid request, should properly update",
				create: &models{
					User: &model.User{
						ID:        123,
						Username:  "username1",
						AvatarURL: "avararurl.com",
						UserStats: json.RawMessage(`{"2023-12-24T12:00:00Z":{"play_count":52,"favorite_count":2, "map_count":3}}`),
						CreatedAt: time.Now().UTC(),
						UpdatedAt: time.Now().UTC(),
					},
					Mapsets: []*model.Mapset{
						{
							ID:          123,
							Artist:      "artist",
							Title:       "title",
							Covers:      json.RawMessage(`{"cover1":"cover1","cover2":"cover2"}`),
							Status:      "graveyard",
							LastUpdated: time.Now().UTC(),
							UserID:      123,
							Creator:     "username1",
							PreviewURL:  "avararurl.com",
							Tags:        "tags tags",
							BPM:         150,
							MapsetStats: json.RawMessage(`{"2023-12-24T12:00:00Z":{"play_count":52,"favorite_count":2}}`),
							CreatedAt:   time.Now().UTC(),
							UpdatedAt:   time.Now().UTC(),
						},
					},
					Beatmaps: []*model.Beatmap{
						{
							ID:               77,
							MapsetID:         123,
							DifficultyRating: 5.3,
							Version:          "version1",
							Accuracy:         7.3,
							AR:               9,
							BPM:              150.3,
							CS:               4,
							Status:           "graveyard",
							URL:              "beatmap1url.com",
							TotalLength:      23,
							UserID:           123,
							LastUpdated:      time.Now().UTC(),
							BeatmapStats:     json.RawMessage(`{"2023-12-24T12:00:00Z":{"play_count":25,"pass_count":23}}`),
							CreatedAt:        time.Now().UTC(),
							UpdatedAt:        time.Now().UTC(),
						},
						{
							ID:               78,
							MapsetID:         123,
							DifficultyRating: 6.7,
							Version:          "version2",
							Accuracy:         8.5,
							AR:               9.3,
							BPM:              150.2,
							CS:               3.3,
							Status:           "graveyard",
							URL:              "beatmap2url.com",
							TotalLength:      24,
							UserID:           123,
							LastUpdated:      time.Now().UTC(),
							BeatmapStats:     json.RawMessage(`{"2023-12-24T12:00:00Z":{"play_count":27,"pass_count":24}}`),
							CreatedAt:        time.Now().UTC(),
							UpdatedAt:        time.Now().UTC(),
						},
					},
				},
				in: &command.UpdateUserCardCommand{
					User: &command.UpdateUserCommand{
						ID:        123,
						AvatarURL: "avatarurlchanged.com",
						Username:  "username1changed",
					},
					Mapsets: []*command.UpdateMapsetCommand{
						{
							Id:     123,
							Artist: "artistchanged",
							Title:  "titlechanged",
							Covers: map[string]string{
								"cover1changed": "cover1changed",
								"cover2changed": "cover2changed",
							},
							Status:        "statuschanged",
							LastUpdated:   time.Now().UTC(),
							UserId:        123,
							PreviewUrl:    "previewurlchanged.com",
							Tags:          "tagschanged tagschanged",
							PlayCount:     200,
							FavoriteCount: 200,
							Bpm:           200,
							Creator:       "username1changed",
							Beatmaps: []*command.UpdateBeatmapCommand{
								{
									Id:               77,
									BeatmapsetId:     123,
									DifficultyRating: 7.6,
									Version:          "version1changed",
									Accuracy:         1,
									Ar:               1,
									Bpm:              1,
									Cs:               1,
									Status:           "graveyard",
									Url:              "urlchanged.com",
									TotalLength:      1,
									UserId:           123,
									Passcount:        100,
									Playcount:        100,
									LastUpdated:      time.Now().UTC(),
								},
								{
									Id:               78,
									BeatmapsetId:     123,
									DifficultyRating: 1.2,
									Version:          "version2changed",
									Accuracy:         2,
									Ar:               2,
									Bpm:              2,
									Cs:               2,
									Status:           "graveyard",
									Url:              "urlchanged.com",
									TotalLength:      1,
									UserId:           123,
									Passcount:        100,
									Playcount:        100,
									LastUpdated:      time.Now().UTC(),
								},
							},
						},
						{
							Id:     345,
							Artist: "artist",
							Title:  "title",
							Covers: map[string]string{
								"cover1": "cover1",
								"cover2": "cover2",
							},
							Status:        "graveyard",
							LastUpdated:   time.Now().UTC(),
							UserId:        123,
							PreviewUrl:    "previewurlnewmap.com",
							Tags:          "tags tags",
							PlayCount:     345,
							FavoriteCount: 456,
							Bpm:           120,
							Creator:       "username1changed",
							Beatmaps: []*command.UpdateBeatmapCommand{
								{
									Id:               1488,
									BeatmapsetId:     345,
									DifficultyRating: 1,
									Version:          "version1",
									Accuracy:         2,
									Ar:               3,
									Bpm:              4,
									Cs:               5,
									Status:           "graveyard",
									Url:              "url.com",
									TotalLength:      1,
									UserId:           123,
									Passcount:        3,
									Playcount:        4,
									LastUpdated:      time.Now().UTC(),
								},
								{
									Id:               1337,
									BeatmapsetId:     345,
									DifficultyRating: 1,
									Version:          "version2",
									Accuracy:         3,
									Ar:               4,
									Bpm:              5,
									Cs:               6,
									Status:           "graveyard",
									Url:              "url.com",
									TotalLength:      1,
									UserId:           123,
									Passcount:        3,
									Playcount:        4,
									LastUpdated:      time.Now().UTC(),
								},
							},
						},
					},
				},
				outCode: 200,
				result: &models{
					User: &model.User{
						ID:        123,
						Username:  "username1changed",
						AvatarURL: "avatarurlchanged.com",
					},
					Mapsets: []*model.Mapset{
						{
							ID:          123,
							Artist:      "artistchanged",
							Title:       "titlechanged",
							Covers:      json.RawMessage(`{"cover1changed":"cover1changed","cover2changed":"cover2changed"}`),
							Status:      "statuschanged",
							LastUpdated: time.Now().UTC(),
							UserID:      123,
							Creator:     "username1changed",
							PreviewURL:  "previewurlchanged.com",
							Tags:        "tagschanged tagschanged",
							BPM:         200,
						},
						{
							ID:          345,
							Artist:      "artist",
							Title:       "title",
							Covers:      json.RawMessage(`{"cover1changed":"cover1changed","cover2changed":"cover2changed"}`),
							Status:      "graveyard",
							LastUpdated: time.Now().UTC(),
							UserID:      123,
							Creator:     "username1changed",
							PreviewURL:  "previewurlnewmap.com",
							Tags:        "tags tags",
							BPM:         120,
						},
					},
					Beatmaps: []*model.Beatmap{
						{
							ID:               77,
							MapsetID:         123,
							DifficultyRating: 7.6,
							Version:          "version1changed",
							Accuracy:         1,
							AR:               1,
							BPM:              1,
							CS:               1,
							Status:           "graveyard",
							URL:              "urlchanged.com",
							TotalLength:      1,
							UserID:           123,
						},
						{
							ID:               78,
							MapsetID:         123,
							DifficultyRating: 1.2,
							Version:          "version2changed",
							Accuracy:         2,
							AR:               2,
							BPM:              2,
							CS:               2,
							Status:           "graveyard",
							URL:              "urlchanged.com",
							TotalLength:      1,
							UserID:           123,
						},
						{
							ID:               1488,
							MapsetID:         345,
							DifficultyRating: 1,
							Version:          "version1",
							Accuracy:         2,
							AR:               3,
							BPM:              4,
							CS:               5,
							Status:           "graveyard",
							URL:              "url.com",
							TotalLength:      1,
							UserID:           123,
						},
						{
							ID:               1337,
							MapsetID:         345,
							DifficultyRating: 1,
							Version:          "version2",
							Accuracy:         3,
							AR:               4,
							BPM:              5,
							CS:               6,
							Status:           "graveyard",
							URL:              "url.com",
							TotalLength:      1,
							UserID:           123,
						},
					},
				},
			},
		}
		for _, tc := range tt {
			s.Run(tc.name, func() {
				// create models for update
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

				inJSON, err := json.Marshal(tc.in)
				s.Require().NoError(err)

				out, err := http.Post(
					fmt.Sprintf("http://localhost:%s/api/user_card/update", s.port),
					"application/json",
					bytes.NewBuffer(inJSON),
				)
				defer out.Body.Close()

				s.Require().NoError(err)
				equal := s.Equal(tc.outCode, out.StatusCode)
				if !equal {
					s.Failf("unexpected response code", "expected %d, got %d", tc.outCode, out.StatusCode)
				}

				// assert result with acual stuff in db
				// user
				expectedUser := tc.result.User

				var actualUser model.User
				err = s.db.Table("users").Where("id = ?", expectedUser.ID).First(&actualUser).Error
				s.Require().NoError(err)

				s.Equal(expectedUser.ID, actualUser.ID)
				s.Equal(expectedUser.AvatarURL, actualUser.AvatarURL)
				s.Equal(expectedUser.Username, actualUser.Username)
				s.Positive(actualUser.CreatedAt.Unix()) // todo
				s.Positive(actualUser.UpdatedAt.Unix()) // todo

				var data map[string]any

				err = json.Unmarshal(actualUser.UserStats, &data)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}

				s.Len(data, 2)

				// mapsets
				expectedMapsets := tc.result.Mapsets

				for _, expectedMapset := range expectedMapsets {
					var actualMapset model.Mapset
					err = s.db.Table("mapsets").Where("id = ?", expectedMapset.ID).First(&actualMapset).Error
					s.Require().NoError(err)

					s.Equal(expectedMapset.ID, actualMapset.ID)
					s.Equal(expectedMapset.Artist, actualMapset.Artist)
					s.Equal(expectedMapset.Title, actualMapset.Title)
					s.Equal(expectedMapset.Status, actualMapset.Status)
					s.Equal(expectedMapset.UserID, actualMapset.UserID)
					s.Equal(expectedMapset.PreviewURL, actualMapset.PreviewURL)
					s.Equal(expectedMapset.Tags, actualMapset.Tags)
					s.Assert().Equal(expectedMapset.BPM, actualMapset.BPM)

					s.Positive(actualMapset.CreatedAt.Unix()) // todo
					s.Positive(actualMapset.UpdatedAt.Unix()) // todo

					var data map[string]any

					err := json.Unmarshal(expectedMapset.MapsetStats, &data)
					if err != nil {
						fmt.Println("Error:", err)
						return
					}

					if actualMapset.ID == 123 {
						s.Len(data, 2)
					} else {
						s.Len(data, 1)
					}
				}

				// beatmaps
				expectedBeatmaps := tc.result.Beatmaps

				for _, expectedBeatmap := range expectedBeatmaps {
					var actualBeatmap model.Beatmap
					err = s.db.Table("beatmaps").Where("id = ?", expectedBeatmap.ID).First(&actualBeatmap).Error
					s.Require().NoError(err)

					s.Equal(expectedBeatmap.ID, actualBeatmap.ID)
					s.Equal(expectedBeatmap.MapsetID, actualBeatmap.MapsetID)
					s.Assert().Equal(expectedBeatmap.DifficultyRating, actualBeatmap.DifficultyRating)
					s.Equal(expectedBeatmap.Version, actualBeatmap.Version)
					s.Assert().Equal(expectedBeatmap.Accuracy, actualBeatmap.Accuracy)
					s.Assert().Equal(expectedBeatmap.AR, actualBeatmap.AR)
					s.Assert().Equal(expectedBeatmap.BPM, actualBeatmap.BPM)
					s.Assert().Equal(expectedBeatmap.CS, actualBeatmap.CS)
					s.Equal(expectedBeatmap.Status, actualBeatmap.Status)
					s.Equal(expectedBeatmap.URL, actualBeatmap.URL)
					s.Equal(expectedBeatmap.TotalLength, actualBeatmap.TotalLength)
					s.Equal(expectedBeatmap.UserID, actualBeatmap.UserID)
					s.Equal(expectedBeatmap.LastUpdated, actualBeatmap.LastUpdated)

					s.Positive(actualBeatmap.CreatedAt.Unix()) // todo
					s.Positive(actualBeatmap.UpdatedAt.Unix()) // todo

					var data map[string]any

					err := json.Unmarshal(expectedBeatmap.BeatmapStats, &data)
					if err != nil {
						fmt.Println("Error:", err)
						return
					}

					if actualBeatmap.ID == 3 || actualBeatmap.ID == 4 {
						s.Len(data, 2)
					} else {
						s.Len(data, 1)
					}
				}
			})
		}
	})
	err := integration.ClearTables(s.ctx, s.db)
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *IntegrationSuite) Test_ProvideUserCard() {
	s.T().Skip("this endpoint disabled for now")
	// todo: might test usecase later

	s.Run("valid requests", func() {
		type models struct {
			User     *model.User
			Mapsets  []*model.Mapset
			Beatmaps []*model.Beatmap
		}

		var tt = []struct {
			name    string
			create  *models
			in      string
			out     *dto.UserCard
			outCode int
		}{
			{
				name: "valid request, should properly provide",
				create: &models{
					User: &model.User{
						ID:        1,
						AvatarURL: "avatarurl.com",
						Username:  "username",
						UserStats: json.RawMessage(`{"2023-12-24T12:00:00Z":{"play_count":100,"favorite_count":2, "map_count":1}}`),
					},
					Mapsets: []*model.Mapset{
						{
							ID:          1,
							Artist:      "artist",
							Title:       "title",
							Covers:      json.RawMessage(`{"cover1": "cover1", "cover2": "cover2"}`),
							Status:      "graveyard",
							LastUpdated: time.Now().UTC(),
							UserID:      1,
							Creator:     "username",
							PreviewURL:  "previewurl.com",
							Tags:        "tags shmags",
							BPM:         210,
							MapsetStats: json.RawMessage(`{"2023-12-24T12:00:00Z":{"play_count":100,"favorite_count":2}}`),
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
							BeatmapStats:     json.RawMessage(`{"2023-12-24T12:00:00Z":{"play_count":52,"pass_count":2}}`),
						},
						{
							ID:               2,
							MapsetID:         1,
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
							BeatmapStats:     json.RawMessage(`{"2023-12-24T12:00:00Z":{"play_count":13,"pass_count":2}}`),
						},
					},
				},
				in: "1",
				out: &dto.UserCard{
					User: &dto.User{
						ID:        1,
						AvatarURL: "avatarurl.com",
						Username:  "username",
					},
					Mapsets: []*dto.Mapset{
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

				url := fmt.Sprintf("http://localhost:%s/api/user_card/", s.port)
				out, err := http.Get(url + tc.in)
				s.Require().NoError(err)
				s.Require().Equal(tc.outCode, out.StatusCode)

				defer out.Body.Close() // Ensure the response body is closed

				body, err := io.ReadAll(out.Body)
				s.Require().NoError(err)

				var actual dto.UserCard
				err = json.Unmarshal(body, &actual)
				s.Require().NoError(err)

				expectedUser := tc.out.User

				s.Equal(expectedUser.ID, actual.User.ID)
				s.Equal(expectedUser.AvatarURL, actual.User.AvatarURL)
				s.Equal(expectedUser.Username, actual.User.Username)

				s.Len(actual.User.UserStats, 1)

				s.Len(actual.Mapsets, 1)

				for _, actualMapset := range actual.Mapsets {
					expectedMapset := tc.out.Mapsets[0]

					s.Equal(expectedMapset.Id, actualMapset.Id)
					s.Equal(expectedMapset.Artist, actualMapset.Artist)
					s.Equal(expectedMapset.Title, actualMapset.Title)
					s.Equal(expectedMapset.Covers, actualMapset.Covers)
					s.Equal(expectedMapset.Status, actualMapset.Status)
					s.Equal(expectedMapset.UserId, actualMapset.UserId)
					s.Equal(expectedMapset.PreviewUrl, actualMapset.PreviewUrl)
					s.Equal(expectedMapset.Tags, actualMapset.Tags)
					s.Assert().Equal(expectedMapset.Bpm, actualMapset.Bpm)
					s.Equal(expectedMapset.Creator, actualMapset.Creator)

					s.Len(actualMapset.MapsetStats, 1)

					for i, actualBeatmap := range actualMapset.Beatmaps {
						expectedBeatmap := tc.out.Mapsets[0].Beatmaps[i]

						s.Equal(expectedBeatmap.Id, actualBeatmap.Id)
						s.Equal(expectedBeatmap.BeatmapsetId, actualBeatmap.BeatmapsetId)
						s.Assert().Equal(expectedBeatmap.DifficultyRating, actualBeatmap.DifficultyRating)
						s.Equal(expectedBeatmap.Version, actualBeatmap.Version)
						s.Assert().Equal(expectedBeatmap.Accuracy, actualBeatmap.Accuracy)
						s.Assert().Equal(expectedBeatmap.Ar, actualBeatmap.Ar)
						s.Assert().Equal(expectedBeatmap.Bpm, actualBeatmap.Bpm)
						s.Assert().Equal(expectedBeatmap.Cs, actualBeatmap.Cs)
						s.Equal(expectedBeatmap.Status, actualBeatmap.Status)
						s.Equal(expectedBeatmap.Url, actualBeatmap.Url)
						s.Equal(expectedBeatmap.TotalLength, actualBeatmap.TotalLength)
						s.Equal(expectedBeatmap.UserId, actualBeatmap.UserId)

						s.Len(actualBeatmap.BeatmapStats, 1)
					}
				}
			})
		}
	})
	err := integration.ClearTables(s.ctx, s.db)
	if err != nil {
		s.T().Fatal(err)
	}
}
