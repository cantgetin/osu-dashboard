package statisticprovide

import (
	"context"
	"math"
	"playcount-monitor-backend/internal/database/txmanager"
	"sort"
	"strconv"
	"strings"
)

type UserMapStatistics struct {
	Tags      map[string]int `json:"most_popular_tags"`
	Languages map[string]int `json:"most_popular_languages"`
	Genres    map[string]int `json:"most_popular_genres"`
	BPMs      map[string]int `json:"most_popular_bpms"`
	Starrates map[string]int `json:"most_popular_starrates"`
}

func (uc *UseCase) GetForUser(
	ctx context.Context,
	userID int,
) (*UserMapStatistics, error) {
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

	return &UserMapStatistics{
		Tags:      tags,
		Languages: languages,
		Genres:    genres,
		BPMs:      BPMs,
		Starrates: starrates,
	}, nil
}

func roundUpToNearestTen(num int) int {
	rounded := int(math.Ceil(float64(num) / 10.0))
	result := rounded * 10
	return result
}

func roundUpToNearestNum(num int) int {
	rounded := int(math.Ceil(float64(num) / 1.0))
	result := rounded * 1
	return result
}

func getTopNKeys(m map[string]int, n int) []string {
	keys := make([]string, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})

	if len(keys) > n {
		keys = keys[:n]
	}

	return keys
}

func filterMapByKey(originalMap map[string]int, keys []string) map[string]int {
	filteredMap := make(map[string]int)

	for _, key := range keys {
		filteredMap[key] = originalMap[key]
	}

	return filteredMap
}

func top5Values(inputMap map[string]int) map[string]int {
	topKeys := getTopNKeys(inputMap, 5)
	result := filterMapByKey(inputMap, topKeys)

	for k, _ := range result {
		if strings.TrimSpace(k) == "" {
			k = "Unspecified"
		}
	}

	return result
}
