package statisticserviceapi

import (
	"github.com/labstack/echo/v4"
)

func (s *ServiceImpl) GetUserMapStatistics(c echo.Context) error {
	// get user id from request context
	userID := c.Get("userID").(string)
	if userID == "" {
		return echo.ErrBadRequest
	}

	//userStatistics := s.statisticProvider.GetForUser()

	return c.JSON(200, nil)
}

type UserMapStatisticsResponse struct {
	MostPopularTags      map[string]int `json:"most_popular_tags"`
	MostPopularGenres    map[string]int `json:"most_popular_genres"`
	MostPopularLanguages map[string]int `json:"most_popular_languages"`
	MostPopularStarrates map[string]int `json:"most_popular_starrates"`
	MostPopularBpms      map[string]int `json:"most_popular_bpms"`
}
