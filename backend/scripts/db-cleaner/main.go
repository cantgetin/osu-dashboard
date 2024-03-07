package main

import (
	"context"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"playcount-monitor-backend/internal/bootstrap"
	"playcount-monitor-backend/internal/config"
	"playcount-monitor-backend/internal/database/repository"
	"playcount-monitor-backend/internal/database/repository/beatmaprepository"
	"playcount-monitor-backend/internal/database/repository/mapsetrepository"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/repository/userrepository"
	"playcount-monitor-backend/internal/database/txmanager"
	"sort"
	"time"
)

const statsMaxElements = 30

func main() {
	cfg, err := config.LoadConfig(".env") // assume we only use this in production
	if err != nil {
		log.Fatalf("failed to load config, %v", err)
	}

	lg := log.New()

	ctx, _ := context.WithCancel(context.Background())

	if err := cleanDBJSONB(ctx, cfg, lg); err != nil {
		log.Fatalf("failed to start db cleaner, %v", err)
	}

}

func cleanDBJSONB(ctx context.Context, cfg *config.Config, log *log.Logger) error {
	db, err := bootstrap.InitDB(cfg)
	if err != nil {
		return err
	}

	err = bootstrap.ApplyMigrations(db)
	if err != nil {
		return err
	}

	// init user, mapset, beatmap repo (things that have stats jsonb)
	userRepo, err := userrepository.New(cfg, log)
	if err != nil {
		return err
	}

	mapsetRepo, err := mapsetrepository.New(cfg, log)
	if err != nil {
		return err
	}

	beatmapRepo, err := beatmaprepository.New(cfg, log)
	if err != nil {
		return err
	}

	txm := bootstrap.ConnectTxManager("", time.Second, db, nil)

	txErr := txm.ReadWrite(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		// get users, trim and update
		var users []*model.User
		users, err = userRepo.List(ctx, tx)
		if err != nil {
			return err
		}

		for _, user := range users {
			err = removeAllMapEntriesExceptLastN(user.UserStats, statsMaxElements)
			if err != nil {
				return err
			}

			err = userRepo.Update(ctx, tx, user)
			if err != nil {
				return err
			}
		}

		// get mapsets, its beatmaps, trim and update
		var mapsets []*model.Mapset
		mapsets, err = mapsetRepo.List(ctx, tx)
		if err != nil {
			return err
		}

		for _, mapset := range mapsets {
			err = removeAllMapEntriesExceptLastN(mapset.MapsetStats, statsMaxElements)
			if err != nil {
				return err
			}

			err = mapsetRepo.Update(ctx, tx, mapset)

			var beatmaps []*model.Beatmap
			beatmaps, err = beatmapRepo.ListForMapset(ctx, tx, mapset.ID)
			if err != nil {
				return err
			}

			for _, beatmap := range beatmaps {
				err = removeAllMapEntriesExceptLastN(beatmap.BeatmapStats, statsMaxElements)
				if err != nil {
					return err
				}

				err = beatmapRepo.Update(ctx, tx, beatmap)
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

	return nil
}

// todo: unit test that
func removeAllMapEntriesExceptLastN(jsonData repository.JSON, n int) error {
	// Unmarshal JSON data into a map[time.Time]interface{}
	data := make(map[time.Time]interface{})
	if err := convertJSONToTimeMap(jsonData, &data); err != nil {
		return err
	}

	// Convert the map keys to a slice for sorting
	keys := make([]time.Time, 0, len(data))
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
		return err
	}

	// Update the original JSON
	jsonData = make(repository.JSON, len(updatedJSON))
	if err := json.Unmarshal(updatedJSON, &jsonData); err != nil {
		return err
	}

	return nil
}

func convertJSONToTimeMap(jsonData repository.JSON, data *map[time.Time]interface{}) error {
	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(jsonBytes, data); err != nil {
		return err
	}
	return nil
}
