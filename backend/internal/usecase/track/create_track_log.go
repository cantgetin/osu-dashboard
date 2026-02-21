package track

import (
	"context"
	"fmt"
	"osu-dashboard/internal/database/model"
	"osu-dashboard/internal/database/txmanager"
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

	txErr := uc.txm.ReadWrite(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		if err := uc.log.Create(ctx, tx, &model.Log{
			Name:               "Daily tracking for all users",
			Message:            model.LogMessageDailyTrack,
			Service:            "tracker",
			AppVersion:         "v1.0",
			Platform:           "Backend",
			Type:               model.LogTypeRegular,
			APIRequests:        reqs,
			SuccessRatePercent: successRate,
			TrackedAt:          time.Now().UTC(),
			AvgResponseTime:    respTime,
			ElapsedTime:        elapsed,
			TimeSinceLastTrack: timeSinceLast,
		}); err != nil {
			return fmt.Errorf("failed to create log: %w", err)
		}
		return nil
	})
	if txErr != nil {
		return txErr
	}

	return nil
}
