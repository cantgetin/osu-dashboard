package statisticprovide

import (
	"context"
	"playcount-monitor-backend/internal/database/txmanager"
	"playcount-monitor-backend/internal/dto"
	"strconv"
	"strings"
)

func (uc *UseCase) GetForUser(ctx context.Context, userID int) (*dto.UserMapStatistics, error) {
	tags := make(map[string]int)
	languages := make(map[string]int)
	genres := make(map[string]int)
	BPMs := make(map[string]int)
	starrates := make(map[string]int)

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
				tags[tag]++
			}

			languages[mapset.Language]++
			genres[mapset.Genre]++

			bpmStr := strconv.Itoa(roundUpToNearestTen(int(mapset.BPM)))
			BPMs[bpmStr]++
		}

		beatmaps, err := uc.beatmap.ListForMapsets(ctx, tx, mapsetIDs...)
		if err != nil {
			return err
		}
		for _, beatmap := range beatmaps {
			starrateStr := strconv.Itoa(roundUpToNearestNum(int(beatmap.DifficultyRating)))
			starrates[starrateStr]++
		}

		return nil
	})
	if txErr != nil {
		return nil, txErr
	}

	// keep only highest 10 values for each map
	tags = top5Values(tags)
	languages = top5Values(languages)
	genres = top5Values(genres)
	BPMs = top5Values(BPMs)
	starrates = top5Values(starrates)

	return &dto.UserMapStatistics{
		Tags:      tags,
		Languages: languages,
		Genres:    genres,
		BPMs:      BPMs,
		Starrates: starrates,
	}, nil
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
