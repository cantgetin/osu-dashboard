package userprovide

import (
	"osu-dashboard/internal/dto"
	osuapimodels "osu-dashboard/internal/service/osuapi/models"
	"time"
)

func MapOsuApiUserToUserDTO(osuUser *osuapimodels.User) (*dto.User, error) {
	return &dto.User{
		ID:            osuUser.ID,
		AvatarURL:     osuUser.AvatarURL,
		Username:      osuUser.Username,
		Tracking:      false,
		TrackingSince: time.Time{},
		UserStats:     nil,
		UserMapCounts: nil,
	}, nil
}
