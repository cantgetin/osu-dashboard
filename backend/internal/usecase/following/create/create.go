package followingcreate

import (
	"context"
	"fmt"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
	"time"
)

func (uc *UseCase) Create(ctx context.Context, code string) error {
	token, err := uc.osuAPI.ExchangeCodeForToken(code)
	if err != nil {
		return fmt.Errorf("failed to exchange code for access token: %w", err)
	}

	user, err := uc.osuAPI.GetUserInfoByHisToken(token.AccessToken)
	if err != nil {
		return fmt.Errorf("failed to get user info: %w", err)
	}

	following := &model.Following{
		ID:          user.ID,
		Username:    user.Username,
		CreatedAt:   time.Now().UTC(),
		LastFetched: time.Now().Add(time.Hour * -25), // hack
	}
	if following.ID <= 0 || following.Username == "" {
		return fmt.Errorf("user id or username is empty: %w", err)
	}

	txErr := uc.txm.ReadWrite(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		f, err := uc.following.Get(ctx, tx, following.ID)
		if f != nil {
			return nil // assume it's already added
		}

		err = uc.following.Create(ctx, tx, following)
		if err != nil {
			return err
		}

		return nil
	})
	if txErr != nil {
		return txErr
	}

	// start tracking background process
	go func() {
		trackCtx := context.Background()
		if err = uc.trackAndCreateRecord(trackCtx, following); err != nil {
			uc.lg.Printf("failed to track and create record %v", err)
		}
	}()

	return nil
}

func (uc *UseCase) trackAndCreateRecord(ctx context.Context, following *model.Following) error {
	started := time.Now().UTC()

	err := uc.track.TrackSingleFollowing(ctx, following)
	if err != nil {
		return fmt.Errorf("failed to track specific user %v", err)
	}

	stats := uc.osuAPI.GetTransportStats()
	defer uc.osuAPI.ResetStats()

	if err = uc.log.Create(ctx, &model.Log{
		Name:               fmt.Sprintf("Initial tracking for user %s", following.Username),
		Message:            model.LogMessageInitialTrack,
		Service:            "osu-dashboard-api",
		AppVersion:         "v1.0",
		Platform:           "Backend",
		Type:               model.TrackTypeInitial,
		TrackedAt:          time.Now().UTC(),
		ElapsedTime:        time.Since(started),
		APIRequests:        stats.RequestCount,
		SuccessRatePercent: stats.SuccessRate,
		AvgResponseTime:    stats.AvgResponseTime,
	}); err != nil {
		return fmt.Errorf("failed to create log: %v", err)
	}

	return nil
}
