package usercreate

import "time"

type CreateUserCommand struct {
	AvatarURL                        string    `json:"avatar_url"`
	CountryCode                      string    `json:"country_code"`
	ID                               int       `json:"id"`
	IsActive                         bool      `json:"is_active"`
	IsBot                            bool      `json:"is_bot"`
	IsDeleted                        bool      `json:"is_deleted"`
	IsOnline                         bool      `json:"is_online"`
	IsSupporter                      bool      `json:"is_supporter"`
	LastVisit                        time.Time `json:"last_visit"`
	PmFriendsOnly                    bool      `json:"pm_friends_only"`
	ProfileColour                    any       `json:"profile_colour"`
	Username                         string    `json:"username"`
	CoverURL                         string    `json:"cover_url"`
	Discord                          any       `json:"discord"`
	HasSupported                     bool      `json:"has_supported"`
	Interests                        any       `json:"interests"`
	JoinDate                         time.Time `json:"join_date"`
	Location                         any       `json:"location"`
	MaxBlocks                        int       `json:"max_blocks"`
	MaxFriends                       int       `json:"max_friends"`
	Occupation                       any       `json:"occupation"`
	Playmode                         string    `json:"playmode"`
	Playstyle                        []string  `json:"playstyle"`
	PostCount                        int       `json:"post_count"`
	ProfileOrder                     []string  `json:"profile_order"`
	Title                            any       `json:"title"`
	TitleURL                         any       `json:"title_url"`
	Twitter                          string    `json:"twitter"`
	Website                          any       `json:"website"`
	AccountHistory                   []any     `json:"account_history"`
	ActiveTournamentBanner           any       `json:"active_tournament_banner"`
	ActiveTournamentBanners          []any     `json:"active_tournament_banners"`
	Badges                           []any     `json:"badges"`
	BeatmapPlaycountsCount           int       `json:"beatmap_playcounts_count"`
	CommentsCount                    int       `json:"comments_count"`
	FavouriteBeatmapsetCount         int       `json:"favourite_beatmapset_count"`
	FollowerCount                    int       `json:"follower_count"`
	Groups                           []any     `json:"groups"`
	GuestBeatmapsetCount             int       `json:"guest_beatmapset_count"`
	LovedBeatmapsetCount             int       `json:"loved_beatmapset_count"`
	MappingFollowerCount             int       `json:"mapping_follower_count"`
	ScoresBestCount                  int       `json:"scores_best_count"`
	ScoresFirstCount                 int       `json:"scores_first_count"`
	ScoresPinnedCount                int       `json:"scores_pinned_count"`
	ScoresRecentCount                int       `json:"scores_recent_count"`
	RankedAndApprovedBeatmapsetCount int       `json:"ranked_and_approved_beatmapset_count"`
	UnrankedBeatmapsetCount          int       `json:"unranked_beatmapset_count"`
	GraveyardBeatmapsetCount         int       `json:"graveyard_beatmapset_count"`
}
