package statisticprovide

import (
	"math"
	"sort"
	"strings"
)

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
