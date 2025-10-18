package track

import (
	"context"
	"fmt"
	"playcount-monitor-backend/internal/database/repository/model"
	"time"
)

func (uc *UseCase) CreateTrackAndLogRecords(ctx context.Context, startTime time.Time, timeSinceLast time.Duration) error {
	var (
		elapsed       = time.Since(startTime)
		reqs          = uc.osuApi.GetOutgoingRequestCount()
		respTime      = uc.osuApi.AverageResponseTime()
		avgReqsPerMin = float64(reqs) / elapsed.Minutes()
		successRate   = uc.osuApi.SuccessRate()
	)

	uc.lg.Infof("Sent %v requests to api in %v minutes", reqs, elapsed.Minutes())
	uc.lg.Infof("Average requests per minute: %f", avgReqsPerMin)

	if err := uc.CreateTrackRecord(ctx); err != nil {
		return fmt.Errorf("failed to create track record: %v", err)
	}

	if err := uc.log.Create(ctx, &model.Log{
		Name:               "Daily tracking for all users",
		Message:            model.LogMessageDailyTrack,
		Service:            "playcount-tracker",
		AppVersion:         "v1.0",
		Platform:           "Backend",
		Type:               model.TrackTypeRegular,
		APIRequests:        reqs,
		SuccessRatePercent: successRate,
		TrackedAt:          time.Now().UTC(),
		AvgResponseTime:    respTime,
		ElapsedTime:        elapsed,
		TimeSinceLastTrack: timeSinceLast,
	}); err != nil {
		return fmt.Errorf("failed to create log: %v", err)
	}

	return nil
}
