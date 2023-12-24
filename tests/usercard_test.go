package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"playcount-monitor-backend/internal/dto"
	usercardcreate "playcount-monitor-backend/internal/usecase/usercard/create"
	"time"
)

func (s *IntegrationSuite) Test_CreateUseCard() {
	s.Run("valid requests", func() {
		tt := []struct {
			name    string
			in      *usercardcreate.CreateUserCardCommand
			outCode int
		}{
			{
				name: "valid request should create",
				in: &usercardcreate.CreateUserCardCommand{
					User: &dto.CreateUserCommand{
						ID:                       1,
						AvatarURL:                "avatarurl.com",
						Username:                 "username",
						UnrankedBeatmapsetCount:  1,
						GraveyardBeatmapsetCount: 1,
					},
					Mapsets: []*dto.CreateMapsetCommand{
						{
							Id:     1,
							Artist: "artist",
							Title:  "title",
							Covers: map[string]string{
								"cover1": "cover1",
								"cover2": "cover2",
							},
							Status:         "graveyard",
							LastUpdated:    time.Now().UTC(),
							UserId:         1,
							PreviewUrl:     "previewurl.com",
							Tags:           "tags tags",
							PlayCount:      20,
							FavouriteCount: 25,
							Bpm:            150,
							Creator:        "username",
							Beatmaps: []*dto.CreateBeatmapCommand{
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

				out, err := http.Post("http://localhost:8080/user_card", "application/json", bytes.NewBuffer(inJSON))

				s.Require().NoError(err)
				s.Require().NotNil(out)

				s.Assert().Equal(out.StatusCode, tc.outCode)
			})
		}
	})
}
