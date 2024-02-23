package mappers

import (
	"encoding/json"
	"fmt"
	"playcount-monitor-backend/internal/database/repository"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/dto"
	"playcount-monitor-backend/internal/usecase/command"
	"reflect"
	"sort"
	"time"
)

// command -> model

func MapCreateUserCardCommandToUserModel(cmd *command.CreateUserCardCommand) (*model.User, error) {
	// get total playcount, favorites, map count
	var playcount int
	var favorites int
	var mapCount int
	var comments int

	for _, ms := range cmd.Mapsets {
		playcount += ms.PlayCount
		favorites += ms.FavouriteCount
		mapCount++
		comments += ms.CommentsCount
	}

	stats, err := mapUserInfoToStatsJSON(playcount, favorites, mapCount, comments)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:                       cmd.User.ID,
		Username:                 cmd.User.Username,
		AvatarURL:                cmd.User.AvatarURL,
		GraveyardBeatmapsetCount: cmd.User.GraveyardBeatmapsetCount,
		UnrankedBeatmapsetCount:  cmd.User.UnrankedBeatmapsetCount,
		UserStats:                stats,
		UpdatedAt:                time.Now().UTC(),
		CreatedAt:                time.Now().UTC(),
	}, nil
}

func MapUpdateUserCardCommandToUserModel(cmd *command.UpdateUserCardCommand) *model.User {
	// get total playcount, favorites, map count
	var playcount int
	var favorites int
	var mapCount int
	var comments int

	for _, ms := range cmd.Mapsets {
		comments += ms.CommentsCount
		playcount += ms.PlayCount
		favorites += ms.FavouriteCount
		mapCount++
	}

	stats, err := mapUserInfoToStatsJSON(playcount, favorites, mapCount, comments)
	if err != nil {
		return nil
	}

	return &model.User{
		ID:                       cmd.User.ID,
		Username:                 cmd.User.Username,
		AvatarURL:                cmd.User.AvatarURL,
		GraveyardBeatmapsetCount: cmd.User.GraveyardBeatmapsetCount,
		UnrankedBeatmapsetCount:  cmd.User.UnrankedBeatmapsetCount,
		UpdatedAt:                time.Now().UTC(),
		UserStats:                stats,
	}
}

func MapCreateMapsetCommandToMapsetModel(mapset *command.CreateMapsetCommand) (*model.Mapset, error) {
	mapsetStats, err := MapMapsetInfoToStatsJSON(mapset.PlayCount, mapset.FavouriteCount, mapset.CommentsCount)
	if err != nil {
		return nil, err
	}

	covers, err := MapMapsetCoversToCoversJSON(mapset.Covers)
	if err != nil {
		return nil, err
	}

	return &model.Mapset{
		ID:            mapset.Id,
		Artist:        mapset.Artist,
		Title:         mapset.Title,
		Covers:        covers,
		Status:        mapset.Status,
		LastUpdated:   mapset.LastUpdated,
		UserID:        mapset.UserId,
		Creator:       mapset.Creator,
		PreviewURL:    mapset.PreviewUrl,
		Tags:          mapset.Tags,
		MapsetStats:   mapsetStats,
		BPM:           mapset.Bpm,
		LastPlaycount: mapset.PlayCount,
		UpdatedAt:     time.Now().UTC(),
		CreatedAt:     time.Now().UTC(),
	}, nil
}

func MapUpdateMapsetCommandToMapsetModel(mapset *command.UpdateMapsetCommand) (*model.Mapset, error) {
	mapsetStats, err := MapMapsetInfoToStatsJSON(mapset.PlayCount, mapset.FavouriteCount, mapset.CommentsCount)
	if err != nil {
		return nil, err
	}

	covers, err := MapMapsetCoversToCoversJSON(mapset.Covers)
	if err != nil {
		return nil, err
	}

	return &model.Mapset{
		ID:            mapset.Id,
		Artist:        mapset.Artist,
		Title:         mapset.Title,
		Covers:        covers,
		Status:        mapset.Status,
		LastUpdated:   mapset.LastUpdated,
		UserID:        mapset.UserId,
		Creator:       mapset.Creator,
		PreviewURL:    mapset.PreviewUrl,
		Tags:          mapset.Tags,
		MapsetStats:   mapsetStats,
		BPM:           mapset.Bpm,
		LastPlaycount: mapset.PlayCount,
		UpdatedAt:     time.Now().UTC(),
	}, nil
}

func MapCreateBeatmapCommandToBeatmapModel(beatmap *command.CreateBeatmapCommand) (*model.Beatmap, error) {
	stats, err := MapBeatmapInfoToStatsJSON(beatmap.Playcount, beatmap.Passcount)
	if err != nil {
		return nil, err
	}

	return &model.Beatmap{
		ID:               beatmap.Id,
		MapsetID:         beatmap.BeatmapsetId,
		DifficultyRating: beatmap.DifficultyRating,
		Version:          beatmap.Version,
		Accuracy:         beatmap.Accuracy,
		AR:               beatmap.Ar,
		BPM:              beatmap.Bpm,
		CS:               beatmap.Cs,
		Status:           beatmap.Status,
		URL:              beatmap.Url,
		TotalLength:      beatmap.TotalLength,
		UserID:           beatmap.UserId,
		LastUpdated:      beatmap.LastUpdated,
		BeatmapStats:     stats,
		UpdatedAt:        time.Now().UTC(),
		CreatedAt:        time.Now().UTC(),
	}, nil
}

func MapUpdateBeatmapCommandToBeatmapModel(beatmap *command.UpdateBeatmapCommand) (*model.Beatmap, error) {
	stats, err := MapBeatmapInfoToStatsJSON(beatmap.Playcount, beatmap.Passcount)
	if err != nil {
		return nil, err
	}

	return &model.Beatmap{
		ID:               beatmap.Id,
		MapsetID:         beatmap.BeatmapsetId,
		DifficultyRating: beatmap.DifficultyRating,
		Version:          beatmap.Version,
		Accuracy:         beatmap.Accuracy,
		AR:               beatmap.Ar,
		BPM:              beatmap.Bpm,
		CS:               beatmap.Cs,
		Status:           beatmap.Status,
		URL:              beatmap.Url,
		TotalLength:      beatmap.TotalLength,
		UserID:           beatmap.UserId,
		LastUpdated:      beatmap.LastUpdated,
		BeatmapStats:     stats,
		UpdatedAt:        time.Now().UTC(),
	}, nil
}

// model -> dto

func MapUserModelsToUserDTOs(users []*model.User) ([]*dto.User, error) {
	res := make([]*dto.User, len(users))
	for i, user := range users {
		var err error
		res[i], err = MapUserModelToUserDTO(user)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func MapUserModelToUserDTO(user *model.User) (*dto.User, error) {
	stats, err := MapStatsJSONToUserStats(user.UserStats)
	if err != nil {
		return nil, err
	}

	return &dto.User{
		ID:            user.ID,
		Username:      user.Username,
		AvatarURL:     user.AvatarURL,
		TrackingSince: user.CreatedAt,
		UserStats:     stats,
	}, nil
}

func MapMapsetModelToMapsetDTO(mapset *model.Mapset, beatmaps []*model.Beatmap) (*dto.Mapset, error) {
	beatmapsDTOs, err := MapBeatmapModelsToBeatmapDTOs(beatmaps)

	covers, err := MapCoversJSONToMapsetCovers(mapset.Covers)
	if err != nil {
		return nil, err
	}

	stats, err := MapStatsJSONToMapsetStats(mapset.MapsetStats)
	if err != nil {
		return nil, err
	}

	return &dto.Mapset{
		Id:          mapset.ID,
		Artist:      mapset.Artist,
		Title:       mapset.Title,
		Covers:      covers,
		Status:      mapset.Status,
		LastUpdated: mapset.LastUpdated,
		UserId:      mapset.UserID,
		PreviewUrl:  mapset.PreviewURL,
		Tags:        mapset.Tags,
		Creator:     mapset.Creator,
		Bpm:         mapset.BPM,
		MapsetStats: stats,
		Beatmaps:    beatmapsDTOs,
	}, nil
}

func MapBeatmapModelsToBeatmapDTOs(beatmaps []*model.Beatmap) ([]*dto.Beatmap, error) {
	res := make([]*dto.Beatmap, len(beatmaps))
	for i, beatmap := range beatmaps {
		var err error
		res[i], err = MapBeatmapModelToBeatmapDTO(beatmap)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func MapBeatmapModelToBeatmapDTO(beatmap *model.Beatmap) (*dto.Beatmap, error) {
	stats, err := MapStatsJSONToBeatmapStats(beatmap.BeatmapStats)
	if err != nil {
		return nil, err
	}

	return &dto.Beatmap{
		Id:               beatmap.ID,
		BeatmapsetId:     beatmap.MapsetID,
		DifficultyRating: beatmap.DifficultyRating,
		Version:          beatmap.Version,
		Accuracy:         beatmap.Accuracy,
		Ar:               beatmap.AR,
		Bpm:              beatmap.BPM,
		Cs:               beatmap.CS,
		Status:           beatmap.Status,
		Url:              beatmap.URL,
		TotalLength:      beatmap.TotalLength,
		UserId:           beatmap.UserID,
		LastUpdated:      beatmap.LastUpdated,
		BeatmapStats:     stats,
	}, nil
}

// covers

func MapMapsetCoversToCoversJSON(m map[string]string) (repository.JSON, error) {
	coversJson, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal mapset covers to json: %w", err)
	}

	return coversJson, nil
}

func MapCoversJSONToMapsetCovers(covers repository.JSON) (map[string]string, error) {
	mapsetCovers := make(map[string]string)
	if err := json.Unmarshal(covers, &mapsetCovers); err != nil {
		return nil, fmt.Errorf("failed to unmarshal mapset covers: %w", err)
	}

	return mapsetCovers, nil
}

// stats -> JSON

func mapUserInfoToStatsJSON(playcount, favorites, mapcount, comments int) (repository.JSON, error) {
	var stats = make(model.UserStats)
	stats[time.Now().UTC()] = &model.UserStatsModel{
		PlayCount: playcount,
		Favorites: favorites,
		MapCount:  mapcount,
		Comments:  comments,
	}

	statsJson, err := json.Marshal(stats)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal user stats: %w", err)
	}

	return statsJson, nil
}

func MapBeatmapInfoToStatsJSON(playcount, passcount int) (repository.JSON, error) {
	var stats = make(model.BeatmapStats)
	stats[time.Now().UTC()] = &model.BeatmapStatsModel{
		Playcount: playcount,
		Passcount: passcount,
	}

	statsJson, err := json.Marshal(stats)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal beatmap stats: %w", err)
	}

	return statsJson, nil
}

func MapMapsetInfoToStatsJSON(playcount, favourites, comments int) (repository.JSON, error) {
	var stats = make(model.MapsetStats)
	stats[time.Now().UTC()] = &model.MapsetStatsModel{
		Playcount: playcount,
		Favorites: favourites,
		Comments:  comments,
	}

	statsJson, err := json.Marshal(stats)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal mapset stats to json: %w", err)
	}

	return statsJson, nil
}

// JSON -> stats
// todo: generic

func MapStatsJSONToUserStats(j repository.JSON) (model.UserStats, error) {
	userStats := make(model.UserStats)
	if err := json.Unmarshal(j, &userStats); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user stats: %w", err)
	}

	return userStats, nil
}

func MapStatsJSONToBeatmapStats(j repository.JSON) (model.BeatmapStats, error) {
	beatmapStats := make(model.BeatmapStats)
	if err := json.Unmarshal(j, &beatmapStats); err != nil {
		return nil, fmt.Errorf("failed to unmarshal beatmap stats: %w", err)
	}

	return beatmapStats, nil
}

func MapStatsJSONToMapsetStats(j repository.JSON) (model.MapsetStats, error) {
	mapsetStats := make(model.MapsetStats)
	if err := json.Unmarshal(j, &mapsetStats); err != nil {
		return nil, fmt.Errorf("failed to unmarshal mapset stats: %w", err)
	}

	return mapsetStats, nil
}

func KeepLastNKeyValuesFromStats(m interface{}, n int) {
	mapValue := reflect.ValueOf(m)
	if mapValue.Kind() != reflect.Map {
		fmt.Println("Not a map")
		return
	}

	keys := mapValue.MapKeys()
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].Interface().(time.Time).Before(keys[j].Interface().(time.Time))
	})

	if len(keys) <= n {
		return
	}

	for i := 0; i < len(keys)-n; i++ {
		mapValue.SetMapIndex(keys[i], reflect.Value{})
	}
}

// JSON -> JSON append
// todo: generic

func AppendNewUserStats(json1, json2 repository.JSON) (repository.JSON, error) {
	// merge two JSONs that are map[time.Time]model.UserStatsModel
	map1 := make(model.UserStats)
	map2 := make(model.UserStats)

	if err := json.Unmarshal(json1, &map1); err != nil {
		return repository.JSON{}, err
	}

	if err := json.Unmarshal(json2, &map2); err != nil {
		return repository.JSON{}, err
	}

	for key, value := range map2 {
		map1[key] = value
	}

	mergedJSON, err := json.Marshal(map1)
	if err != nil {
		return repository.JSON{}, err
	}

	return mergedJSON, nil
}

func AppendNewMapsetStats(json1, json2 repository.JSON) (repository.JSON, error) {
	// merge two JSONs that are map[time.Time]model.MapsetStatsModel
	map1 := make(model.MapsetStats)
	map2 := make(model.MapsetStats)

	if err := json.Unmarshal(json1, &map1); err != nil {
		return repository.JSON{}, err
	}

	if err := json.Unmarshal(json2, &map2); err != nil {
		return repository.JSON{}, err
	}

	for key, value := range map2 {
		map1[key] = value
	}

	mergedJSON, err := json.Marshal(map1)
	if err != nil {
		return repository.JSON{}, err
	}

	return mergedJSON, nil
}

func AppendNewBeatmapStats(json1, json2 repository.JSON) (repository.JSON, error) {
	// merge two JSONs that are map[time.Time]model.BeatmapStatsModel
	map1 := make(model.BeatmapStats)
	map2 := make(model.BeatmapStats)

	if err := json.Unmarshal(json1, &map1); err != nil {
		return repository.JSON{}, err
	}

	if err := json.Unmarshal(json2, &map2); err != nil {
		return repository.JSON{}, err
	}

	for key, value := range map2 {
		map1[key] = value
	}

	mergedJSON, err := json.Marshal(map1)
	if err != nil {
		return repository.JSON{}, err
	}

	return mergedJSON, nil
}

// other

func MapStatusesToUserMapCounts(statuses []string) *dto.UserMapCounts {
	res := &dto.UserMapCounts{}

	for _, status := range statuses {
		switch status {
		case "graveyard":
			res.Graveyard++
		case "wip":
			res.Wip++
		case "pending":
			res.Pending++
		case "ranked":
			res.Ranked++
		case "approved":
			res.Approved++
		case "qualified":
			res.Qualified++
		case "loved":
			res.Loved++
		}
	}

	return res
}
