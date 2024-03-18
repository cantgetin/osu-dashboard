package mapsetserviceapi

import "playcount-monitor-backend/internal/dto"

type MapsetListResponse struct {
	Mapsets     []*dto.Mapset `json:"mapsets"`
	CurrentPage int           `json:"current_page"`
	Pages       int           `json:"pages"`
}
