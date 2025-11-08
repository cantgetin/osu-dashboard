package userprovide

import (
	"osu-dashboard/internal/dto"
	"osu-dashboard/internal/service/osuapi"
	"time"
)

func MapOsuApiUserToUserDTO(osuUser *osuapi.User) (*dto.User, error) {
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
