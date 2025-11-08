package cleaner

import (
	"context"
	"encoding/json"
	"fmt"
	"osu-dashboard/internal/database/repository"
	"osu-dashboard/internal/database/repository/model"
	"osu-dashboard/internal/database/txmanager"
	"sort"
	"time"
)

const jsonbStatsMaxElements = 14

func (uc *UseCase) Clean(ctx context.Context) error {
	started := time.Now()

	txErr := uc.txm.ReadWrite(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		// get users, trim and update
		users, err := uc.user.List(ctx, tx)
		if err != nil {
			return err
		}

		for _, user := range users {
			updatedStats, err := removeAllMapEntriesExceptLastN(&user.UserStats, jsonbStatsMaxElements)
			if err != nil {
				return err
			}

			user.UserStats = *updatedStats
			err = uc.user.Update(ctx, tx, user)
			if err != nil {
				return err
			}
		}

		// get mapsets, its beatmaps, trim and update
		mapsets, err := uc.mapset.List(ctx, tx)
		if err != nil {
			return err
		}

		for _, mapset := range mapsets {
			updatedMapsetStats, err := removeAllMapEntriesExceptLastN(&mapset.MapsetStats, jsonbStatsMaxElements)
			if err != nil {
				return err
			}

			mapset.MapsetStats = *updatedMapsetStats
			err = uc.mapset.Update(ctx, tx, mapset)
			if err != nil {
				return err
			}

			beatmaps, err := uc.beatmap.ListForMapset(ctx, tx, mapset.ID)
			if err != nil {
				return err
			}

			for _, beatmap := range beatmaps {
				beatmapStats, err := removeAllMapEntriesExceptLastN(&beatmap.BeatmapStats, jsonbStatsMaxElements)
				if err != nil {
					return err
				}

				beatmap.BeatmapStats = *beatmapStats
				err = uc.beatmap.Update(ctx, tx, beatmap)
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
	if txErr != nil {
		return txErr
	}

	if err := uc.log.Create(ctx, &model.Log{
		Name:               "Daily data clean for all users",
		Message:            model.LogMessageDailyClean,
		Service:            "db-cleaner",
		AppVersion:         "v1.0",
		Platform:           "Backend",
		Type:               model.TrackTypeRegular,
		APIRequests:        0,
		SuccessRatePercent: 100,
		TrackedAt:          time.Now().UTC(),
		AvgResponseTime:    0,
		ElapsedTime:        time.Since(started),
		TimeSinceLastTrack: 0,
	}); err != nil {
		return fmt.Errorf("failed to create log: %v", err)
	}

	return nil
}

func removeAllMapEntriesExceptLastN(jsonData *repository.JSON, n int) (*repository.JSON, error) {
	if jsonData == nil {
		return nil, fmt.Errorf("jsonData is nil")
	}

	if *jsonData == nil {
		return nil, fmt.Errorf("jsonData ptr is nil")
	}

	var data map[time.Time]interface{}
	if err := json.Unmarshal(*jsonData, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	if len(data) <= n {
		return jsonData, nil
	}

	// Convert the map keys to a slice for sorting
	keys := make([]time.Time, len(data))
	for key := range data {
		keys = append(keys, key)
	}

	// Sort the keys in ascending order
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].Before(keys[j])
	})

	// Remove all entries except the last n
	for i := 0; i < len(keys)-n; i++ {
		delete(data, keys[i])
	}

	// Marshal the updated map back to JSON
	updatedJSON, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal updated JSON: %v", err)
	}

	r := repository.JSON(updatedJSON)
	return &r, nil
}
