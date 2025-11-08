package dtosort

import (
	"osu-dashboard/internal/dto"
	"time"
)

type MapsetByLastKeyValue []*dto.Mapset

func (a MapsetByLastKeyValue) Len() int      { return len(a) }
func (a MapsetByLastKeyValue) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a MapsetByLastKeyValue) Less(i, j int) bool {
	// Get the last key-value pair for struct i
	var iLastTime time.Time
	var iLastValue int
	for t, v := range a[i].MapsetStats {
		if t.After(iLastTime) {
			iLastTime = t
			iLastValue = v.Playcount
		}
	}

	// Get the last key-value pair for struct j
	var jLastTime time.Time
	var jLastValue int
	for t, v := range a[j].MapsetStats {
		if t.After(jLastTime) {
			jLastTime = t
			jLastValue = v.Playcount
		}
	}

	// Compare last values, sorting in descending order
	if iLastValue != jLastValue {
		return iLastValue > jLastValue
	}

	// If last values are the same, compare by last time
	return iLastTime.After(jLastTime)
}
