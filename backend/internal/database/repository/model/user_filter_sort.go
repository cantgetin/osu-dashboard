package model

// filter

type UserFilterField string

const (
	UserNameField UserFilterField = "username"
)

type UserFilter map[UserFilterField]any

// sort

type UserMapStatsSortFields string

const (
	UserPlaycount UserMapStatsSortFields = "play_count"
	UserMapCount  UserMapStatsSortFields = "map_count"
	UserFavs      UserMapStatsSortFields = "comments_count"
	UserComms     UserMapStatsSortFields = "favorite_count"
)

type UserSort struct {
	Field     UserMapStatsSortFields
	Direction SortDirection
}
