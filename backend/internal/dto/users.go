package dto

type UsersPaged struct {
	Users       []*User `json:"users,omitempty"`
	CurrentPage int     `json:"current_page,omitempty"`
	Pages       int     `json:"pages,omitempty"`
}
