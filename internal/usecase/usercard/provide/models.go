package usercardprovide

import "playcount-monitor-backend/internal/dto"

type UserCard struct {
	User    *dto.User
	Mapsets []*dto.Mapset
}
