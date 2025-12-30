package cleaner

import (
	"context"
	"fmt"
	"osu-dashboard/internal/database/repository"
	"osu-dashboard/internal/database/repository/model"
	"osu-dashboard/internal/database/txmanager"
	"time"
)

const jsonbStatsMaxElements = 14

func (uc *UseCase) Clean(ctx context.Context) error {
	started := time.Now()

	// list users, trim stats and update
	txErr := uc.txm.ReadWrite(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		users, err := uc.user.List(ctx, tx)
		if err != nil {
			return err
		}

		for _, user := range users {
			var updatedStats *repository.JSON
			updatedStats, err = removeAllMapEntriesExceptLastN(&user.UserStats, jsonbStatsMaxElements)
			if err != nil {
				return err
			}

			user.UserStats = *updatedStats
			err = uc.user.Update(ctx, tx, user)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if txErr != nil {
		return fmt.Errorf("failed to clean users: %w", txErr)
	}

	// list mapsets, its beatmaps, trim and update
	txErr = uc.txm.ReadWrite(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		mapsets, err := uc.mapset.List(ctx, tx)
		if err != nil {
			return err
		}

		for _, mapset := range mapsets {
			var updatedMapsetStats *repository.JSON
			updatedMapsetStats, err = removeAllMapEntriesExceptLastN(&mapset.MapsetStats, jsonbStatsMaxElements)
			if err != nil {
				return err
			}

			mapset.MapsetStats = *updatedMapsetStats
			err = uc.mapset.Update(ctx, tx, mapset)
			if err != nil {
				return err
			}

			var beatmaps []*model.Beatmap
			beatmaps, err = uc.beatmap.ListForMapset(ctx, tx, mapset.ID)
			if err != nil {
				return err
			}

			for _, beatmap := range beatmaps {
				var beatmapStats *repository.JSON
				beatmapStats, err = removeAllMapEntriesExceptLastN(&beatmap.BeatmapStats, jsonbStatsMaxElements)
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
		return fmt.Errorf("failed to clean mapsets and beatmaps: %w", txErr)
	}

	// create log
	txErr = uc.txm.ReadWrite(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		if err := uc.log.Create(ctx, tx, &model.Log{
			Name:               "Daily data clean for all users",
			Message:            model.LogMessageDailyClean,
			Service:            "db-cleaner",
			AppVersion:         "v1.0",
			Platform:           "Backend",
			Type:               model.LogTypeRegular,
			APIRequests:        0,
			SuccessRatePercent: 100,
			TrackedAt:          time.Now().UTC(),
			AvgResponseTime:    0,
			ElapsedTime:        time.Since(started),
			TimeSinceLastTrack: 0,
		}); err != nil {
			return err
		}
		return nil
	})
	if txErr != nil {
		return fmt.Errorf("failed to create clean log: %w", txErr)
	}

	return nil
}

func (uc *UseCase) CreateCleanRecord(ctx context.Context) error {
	track := &model.Clean{
		CleanedAt: time.Now().UTC(),
	}

	if err := uc.txm.ReadWrite(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		err := uc.clean.Create(ctx, tx, track)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) GetLastTimeCleaned(
	ctx context.Context,
) (*time.Time, error) {
	t := time.Time{}
	if err := uc.txm.ReadOnly(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		track, err := uc.clean.GetLastClean(ctx, tx)
		if err != nil {
			return err
		}

		t = track.CleanedAt

		return nil
	}); err != nil {
		return nil, err
	}

	return &t, nil
}
