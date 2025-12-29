package mapsetrepository

import (
	"osu-dashboard/internal/database/repository/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_buildListByFilterQuery(t *testing.T) {
	tt := []struct {
		name           string
		filter         model.MapsetFilter
		expectedQuery  string
		expectedValues []any
	}{
		{
			name: "Artist and Title and Tags",
			filter: model.MapsetFilter{
				model.MapsetArtistField: "Artist",
				model.MapsetTitleField:  "Title",
				model.MapsetTagsField:   "Tags",
			},
			expectedQuery:  "artist = ? AND tags = ? AND title = ?",
			expectedValues: []any{"Artist", "Tags", "Title"},
		},
		{
			name: "Artist or Title or Tags",
			filter: model.MapsetFilter{
				model.MapsetArtistOrTitleOrTagsFields: "Search",
			},
			expectedQuery:  "( artist ILIKE ? OR title ILIKE ? OR tags ILIKE ? )",
			expectedValues: []any{"%Search%", "%Search%", "%Search%"},
		},
		{
			name: "Artist or Title or Tags and Status",
			filter: model.MapsetFilter{
				model.MapsetArtistOrTitleOrTagsFields: "Search",
				model.MapsetStatusField:               "Status",
			},
			expectedQuery:  "( artist ILIKE ? OR title ILIKE ? OR tags ILIKE ? ) AND status = ?",
			expectedValues: []any{"%Search%", "%Search%", "%Search%", "Status"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			query, values := buildListByFilterQuery(tc.filter)
			assert.Equal(t, tc.expectedQuery, query)
			assert.Equal(t, tc.expectedValues, values)
		})
	}
}
