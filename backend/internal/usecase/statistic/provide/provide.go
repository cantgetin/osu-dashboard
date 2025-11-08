package statisticprovide

import (
	"context"
	"osu-dashboard/internal/database/txmanager"
	"osu-dashboard/internal/dto"
	"strconv"
	"strings"
)

const ItemsCount = 5

func (uc *UseCase) GetForUser(ctx context.Context, userID int) (*dto.UserMapStatistics, error) {
	res := &dto.UserMapStatistics{
		Tags:      make(map[string]int),
		Languages: make(map[string]int),
		Genres:    make(map[string]int),
		BPMs:      make(map[string]int),
		Starrates: make(map[string]int),
		Combined:  make([]string, 0),
	}

	txErr := uc.txm.ReadOnly(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		mapsets, err := uc.mapset.ListForUser(ctx, tx, userID)
		if err != nil {
			return err
		}

		mapsetIDs := make([]int, len(mapsets))

		for i, mapset := range mapsets {
			mapsetIDs[i] = mapset.ID
			tagsArr := strings.Fields(mapset.Tags)
			for _, tag := range tagsArr {
				res.Tags[tag]++
			}

			res.Languages[mapset.Language]++
			res.Genres[mapset.Genre]++

			bpmStr := strconv.Itoa(roundUpToNearestTen(int(mapset.BPM)))
			res.BPMs[bpmStr]++
		}

		beatmaps, err := uc.beatmap.ListForMapsets(ctx, tx, mapsetIDs...)
		if err != nil {
			return err
		}
		for _, beatmap := range beatmaps {
			starrateStr := strconv.Itoa(int(beatmap.DifficultyRating))
			res.Starrates[starrateStr]++
		}

		return nil
	})
	if txErr != nil {
		return nil, txErr
	}

	// format user statistic
	res.Tags = getTopNValues(res.Tags, ItemsCount)
	res.Languages = getTopNValues(res.Languages, ItemsCount)
	res.Genres = getTopNValues(res.Genres, ItemsCount)

	res.BPMs = getTopNValues(res.BPMs, ItemsCount)
	res.BPMs = appendToAllKeys(res.BPMs, " bpm")

	res.Starrates = getTopNValues(res.Starrates, ItemsCount)
	res.Starrates = appendToAllKeys(res.Starrates, " star")

	res.Combined = combineMapKeys(res.Tags, res.Languages, res.Genres)
	res.Combined = append(res.Combined, getTopKey(res.Starrates), getTopKey(res.BPMs))

	return res, nil
}

func (uc *UseCase) GetForSystem(ctx context.Context) (*dto.SystemStatistics, error) {
	res := new(dto.SystemStatistics)
	txErr := uc.txm.ReadOnly(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		var err error
		res.Users, err = uc.user.TotalCount(ctx, tx)
		if err != nil {
			return err
		}

		res.Mapsets, err = uc.mapset.TotalCount(ctx, tx)
		if err != nil {
			return err
		}

		res.Beatmaps, err = uc.beatmap.TotalCount(ctx, tx)
		if err != nil {
			return err
		}

		res.Tracks, err = uc.track.TotalCount(ctx, tx)
		if err != nil {
			return err
		}

		return nil
	})
	if txErr != nil {
		return nil, txErr
	}

	return res, nil
}
