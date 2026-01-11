package dto

type SearchResultType string

const (
	UserResult   SearchResultType = "user"
	MapsetResult SearchResultType = "mapset"
)

type SearchResult struct {
	Title      string           `json:"title"`
	PictureURL string           `json:"picture_url"`
	Type       SearchResultType `json:"type"`
}
