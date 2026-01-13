package osuapimodels

type Country struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type User struct {
	ID        int      `json:"id"`
	AvatarURL string   `json:"avatar_url"`
	Username  string   `json:"username"`
	Country   *Country `json:"country"`
}
