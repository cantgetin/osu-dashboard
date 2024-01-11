package track

import (
	"playcount-monitor-backend/internal/service/osuapi"
	"playcount-monitor-backend/internal/usecase/command"
	"time"
)

func mapOsuAPiUserToCreateUserCommand(user *osuapi.User) *command.CreateUserCommand {
	return &command.CreateUserCommand{
		ID:                       0,
		AvatarURL:                "",
		Username:                 "",
		UnrankedBeatmapsetCount:  0,
		GraveyardBeatmapsetCount: 0,
	}
}

func mapOsuApiMapsetsToCreateMapsetCommands(mapsets []*osuapi.Beatmap) []*command.CreateMapsetCommand {
	cmds := []*command.CreateMapsetCommand{}
	for _, _ = range mapsets {
		cmds = append(cmds, &command.CreateMapsetCommand{
			Id:             0,
			Artist:         "",
			Title:          "",
			Covers:         nil,
			Status:         "",
			LastUpdated:    time.Time{},
			UserId:         0,
			PreviewUrl:     "",
			Tags:           "",
			PlayCount:      0,
			FavouriteCount: 0,
			Bpm:            0,
			Creator:        "",
			Beatmaps:       nil,
		})
	}
	return cmds
}
