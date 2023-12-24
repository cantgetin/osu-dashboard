package usercardupdate

import "playcount-monitor-backend/internal/dto"

type UpdateUserCardCommand struct {
	User    *dto.CreateUserCommand
	Mapsets []*dto.CreateMapsetCommand
}
