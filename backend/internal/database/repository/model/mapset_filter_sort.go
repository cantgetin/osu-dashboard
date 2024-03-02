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

type MapsetFilter map[MapsetFilterField]interface{}

// sort

type MapsetSortField string

const (
	MapsetPlaycount MapsetSortField = "last_playcount"
	MapsetCreatedAt MapsetSortField = "created_at"
	MapsetFavs      MapsetSortField = "last_favorites"
	MapsetComms     MapsetSortField = "last_comments"
)

type SortDirection string

const (
	ASC  SortDirection = "ASC"
	DESC SortDirection = "DESC"
)

type MapsetSort struct {
	Field     MapsetSortField
	Direction SortDirection
}
