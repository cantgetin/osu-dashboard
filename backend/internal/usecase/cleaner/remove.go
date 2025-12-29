package cleaner

import (
	"encoding/json"
	"fmt"
	"osu-dashboard/internal/database/repository"
	"sort"
	"time"
)

func removeAllMapEntriesExceptLastN(jsonData *repository.JSON, n int) (*repository.JSON, error) {
	if jsonData == nil {
		return nil, fmt.Errorf("jsonData is nil")
	}

	if *jsonData == nil {
		return nil, fmt.Errorf("jsonData ptr is nil")
	}

	var data map[time.Time]any
	if err := json.Unmarshal(*jsonData, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	if len(data) <= n {
		return jsonData, nil
	}

	keys := make([]time.Time, len(data))
	for key := range data {
		keys = append(keys, key)
	}

	// sort in ascending order
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].Before(keys[j])
	})

	for i := 0; i < len(keys)-n; i++ {
		delete(data, keys[i])
	}

	updatedJSON, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal updated JSON: %w", err)
	}

	r := repository.JSON(updatedJSON)
	return &r, nil
}
