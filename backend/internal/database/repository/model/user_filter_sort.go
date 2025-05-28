package model

// filter

type UserFilterField string

const (
	UserNameField UserFilterField = "username"
)

type UserFilter map[UserFilterField]interface{}

// sort

type UserSortField string

const (
	UserPlaycount UserSortField = "playcount"
	UserMapCount  UserSortField = "map_count"
	UserFavs      UserSortField = "favourites"
	UserComms     UserSortField = "comments"
)

type UserSort struct {
	Field     UserSortField
	Direction SortDirection
}
