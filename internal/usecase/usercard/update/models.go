package usercardupdate

import "playcount-monitor-backend/internal/dto"

type UpdateUserCardCommand struct {
	User    *dto.User
	Mapsets []*dto.Mapset
}
