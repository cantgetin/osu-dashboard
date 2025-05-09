package statisticprovide

import (
	"sort"
	"strings"
)

const Unspecified = "Unspecified"
const unspecified = "unspecified"

func roundUpToNearestTen(num int) int {
	return ((num + 9) / 10) * 10
}

func getTopNValues(m map[string]int, n int) map[string]int {
	if n <= 0 || len(m) == 0 {
		return make(map[string]int)
	}

	if n > len(m) {
		n = len(m)
	}

	type kv struct {
		key string
		val int
	}
	kvs := make([]kv, 0, len(m))
	for k, v := range m {
		key := strings.TrimSpace(k)
		if key == "" {
			key = Unspecified
		}
		kvs = append(kvs, kv{strings.ToLower(key), v})
	}

	sort.Slice(kvs, func(i, j int) bool { return kvs[i].val > kvs[j].val })

	result := make(map[string]int, n)
	for i := 0; i < n; i++ {
		result[kvs[i].key] = kvs[i].val
	}

	return result
}

func combineMapKeys(maps ...map[string]int) []string {
	unique := make(map[string]struct{}, 32)
	keys := make([]string, 0, 32)

	unique[Unspecified] = struct{}{}
	unique[unspecified] = struct{}{}

	for _, m := range maps {
		for k := range m {
			if _, exists := unique[k]; !exists {
				unique[k] = struct{}{}
				keys = append(keys, k)
			}
		}
	}

	return keys
}

func appendToAllKeys(m map[string]int, s string) map[string]int {
	result := make(map[string]int, len(m))
	for k, v := range m {
		result[k+s] = v
	}
	return result
}

func getTopKey(m map[string]int) string {
	if len(m) == 0 {
		return ""
	}

	var topKey string
	maxValue := -1 << 31

	for k, v := range m {
		if v > maxValue {
			maxValue = v
			topKey = k
		}
	}

	return topKey
}
