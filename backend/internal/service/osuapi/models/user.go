package osuapimodels

type User struct {
	ID          int    `json:"id"`
	AvatarURL   string `json:"avatar_url"`
	Username    string `json:"username"`
	CountryCode string `json:"country_code"`
}
