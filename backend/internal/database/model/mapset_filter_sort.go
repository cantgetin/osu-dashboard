package model

// filter

type MapsetFilterField string

const (
	MapsetStatusField               MapsetFilterField = "status"
	MapsetArtistField               MapsetFilterField = "artist"
	MapsetTitleField                MapsetFilterField = "title"
	MapsetTagsField                 MapsetFilterField = "tags"
	MapsetArtistOrTitleOrTagsFields MapsetFilterField = ""
)

type MapsetFilter map[MapsetFilterField]any

// sort

type MapsetSortField string

const (
	MapsetPlaycount MapsetSortField = "last_playcount"
	MapsetCreatedAt MapsetSortField = "created_at"
	MapsetFavs      MapsetSortField = "last_favorites"
	MapsetComms     MapsetSortField = "last_comments"
)

type MapsetSort struct {
	Field     MapsetSortField
	Direction SortDirection
}
