package dto

type SearchResultType string

const (
	UserResult   SearchResultType = "user"
	MapsetResult SearchResultType = "mapset"
)

type SearchResult struct {
	ID         int              `json:"id"`
	Title      string           `json:"title"`
	PictureURL string           `json:"picture_url"`
	Type       SearchResultType `json:"type"`
}
