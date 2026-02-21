package logprovide

import (
	"osu-dashboard/internal/database/model"
	"osu-dashboard/internal/dto"
)

func mapLogModelsToLogsDTOs(models []*model.Log) []dto.Log {
	res := make([]dto.Log, 0, len(models))

	for _, l := range models {
		if l != nil {
			res = append(res, mapLogModelToLogDTO(l))
		}
	}
	return res
}

func mapLogModelToLogDTO(model *model.Log) dto.Log {
	return dto.Log{
		ID:                 model.ID,
		Name:               model.Name,
		Message:            string(model.Message),
		Service:            model.Service,
		AppVersion:         model.AppVersion,
		Platform:           model.Platform,
		Type:               string(model.Type),
		APIRequests:        model.APIRequests,
		SuccessRatePercent: model.SuccessRatePercent,
		TrackedAt:          model.TrackedAt,
		AvgResponseTime:    model.AvgResponseTime,
		ElapsedTime:        model.ElapsedTime,
		TimeSinceLastTrack: model.TimeSinceLastTrack,
	}
}
