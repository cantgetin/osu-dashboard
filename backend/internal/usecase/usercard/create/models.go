package usercardcreate

import "playcount-monitor-backend/internal/dto"

type CreateUserCardCommand struct {
	User    *dto.CreateUserCommand
	Mapsets []*dto.CreateMapsetCommand
}
