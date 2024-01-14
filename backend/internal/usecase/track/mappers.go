package track

import (
	"playcount-monitor-backend/internal/service/osuapi"
	"playcount-monitor-backend/internal/usecase/command"
)

// create

func mapOsuApiUserToCreateUserCommand(user *osuapi.User) *command.CreateUserCommand {
	return &command.CreateUserCommand{
		ID:                       user.ID,
		Username:                 user.Username,
		AvatarURL:                user.AvatarURL,
		UnrankedBeatmapsetCount:  user.UnrankedBeatmapsetCount,
		GraveyardBeatmapsetCount: user.GraveyardBeatmapsetCount,
	}
}

func mapOsuApiMapsetsToCreateMapsetCommands(mapsets []*osuapi.Mapset) []*command.CreateMapsetCommand {
	cmds := []*command.CreateMapsetCommand{}
	for _, m := range mapsets {
		cmds = append(cmds, &command.CreateMapsetCommand{
			Id:             m.Id,
			Artist:         m.Artist,
			Title:          m.Title,
			Covers:         m.Covers,
			Status:         m.Status,
			LastUpdated:    m.LastUpdated,
			UserId:         m.UserId,
			PreviewUrl:     m.PreviewUrl,
			Tags:           m.Tags,
			PlayCount:      m.PlayCount,
			FavouriteCount: m.FavouriteCount,
			Bpm:            m.Bpm,
			Creator:        m.Creator,
			Beatmaps:       mapOsuApiBeatmapsToCreateBeatmapCommands(m.Beatmaps),
		})
	}
	return cmds
}

func mapOsuApiBeatmapsToCreateBeatmapCommands(beatmaps []*osuapi.Beatmap) []*command.CreateBeatmapCommand {
	cmds := []*command.CreateBeatmapCommand{}
	for _, b := range beatmaps {
		cmds = append(cmds, &command.CreateBeatmapCommand{
			Id:               b.Id,
			BeatmapsetId:     b.BeatmapsetId,
			DifficultyRating: b.DifficultyRating,
			Version:          b.Version,
			Accuracy:         b.Accuracy,
			Ar:               b.Ar,
			Bpm:              b.Bpm,
			Cs:               b.Cs,
			Status:           b.Status,
			Url:              b.Url,
			TotalLength:      b.TotalLength,
			UserId:           b.UserId,
			Passcount:        b.Passcount,
			Playcount:        b.Playcount,
			LastUpdated:      b.LastUpdated,
		})
	}
	return cmds
}

// update

func mapOsuApiUserToUpdateUserCommand(user *osuapi.User) *command.UpdateUserCommand {
	return &command.UpdateUserCommand{
		ID:                       user.ID,
		Username:                 user.Username,
		AvatarURL:                user.AvatarURL,
		UnrankedBeatmapsetCount:  user.UnrankedBeatmapsetCount,
		GraveyardBeatmapsetCount: user.GraveyardBeatmapsetCount,
	}
}

func mapOsuApiMapsetsToUpdateMapsetCommands(mapsets []*osuapi.Mapset) []*command.UpdateMapsetCommand {
	cmds := []*command.UpdateMapsetCommand{}
	for _, m := range mapsets {
		cmds = append(cmds, &command.UpdateMapsetCommand{
			Id:             m.Id,
			Artist:         m.Artist,
			Title:          m.Title,
			Covers:         m.Covers,
			Status:         m.Status,
			LastUpdated:    m.LastUpdated,
			UserId:         m.UserId,
			PreviewUrl:     m.PreviewUrl,
			Tags:           m.Tags,
			PlayCount:      m.PlayCount,
			FavouriteCount: m.FavouriteCount,
			Bpm:            m.Bpm,
			Creator:        m.Creator,
			Beatmaps:       mapOsuApiBeatmapsToUpdateBeatmapCommands(m.Beatmaps),
		})
	}
	return cmds
}

func mapOsuApiBeatmapsToUpdateBeatmapCommands(beatmaps []*osuapi.Beatmap) []*command.UpdateBeatmapCommand {
	cmds := []*command.UpdateBeatmapCommand{}
	for _, b := range beatmaps {
		cmds = append(cmds, &command.UpdateBeatmapCommand{
			Id:               b.Id,
			BeatmapsetId:     b.BeatmapsetId,
			DifficultyRating: b.DifficultyRating,
			Version:          b.Version,
			Accuracy:         b.Accuracy,
			Ar:               b.Ar,
			Bpm:              b.Bpm,
			Cs:               b.Cs,
			Status:           b.Status,
			Url:              b.Url,
			TotalLength:      b.TotalLength,
			UserId:           b.UserId,
			Passcount:        b.Passcount,
			Playcount:        b.Playcount,
			LastUpdated:      b.LastUpdated,
		})
	}

	return cmds
}
