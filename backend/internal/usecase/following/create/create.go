package followingcreate

import (
	"context"
	"fmt"
	"osu-dashboard/internal/database/model"
	"osu-dashboard/internal/database/txmanager"
	"osu-dashboard/internal/dto"
	userprovide "osu-dashboard/internal/usecase/user/provide"
	"time"
)

func (uc *UseCase) Create(ctx context.Context, code string) (*dto.User, error) {
	token, err := uc.osuAPI.ExchangeCodeForToken(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange code for access token: %w", err)
	}

	user, err := uc.osuAPI.GetUserInfoByHisToken(ctx, token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %w", err)
	}

	userDTO, err := userprovide.MapOsuApiUserToUserDTO(user)
	if err != nil {
		return nil, fmt.Errorf("failed to map osu user dto: %w", err)
	}

	following := &model.Following{
		ID:          user.ID,
		Username:    user.Username,
		CreatedAt:   time.Now().UTC(),
		LastFetched: time.Now().Add(time.Hour * -25), // hack
	}
	if following.ID <= 0 || following.Username == "" {
		return nil, fmt.Errorf("user id or username is empty: %w", err)
	}

	txErr := uc.txm.ReadWrite(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		var f *model.Following
		f, err = uc.following.Get(ctx, tx, following.ID)
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
		return nil, txErr
	}

	// start tracking background process
	go func() {
		trackCtx := context.Background()
		if err = uc.trackAndCreateRecord(trackCtx, following); err != nil {
			uc.lg.Printf("failed to track and create record %v", err)
		}
	}()

	return userDTO, nil
}

func (uc *UseCase) trackAndCreateRecord(ctx context.Context, following *model.Following) error {
	started := time.Now().UTC()

	err := uc.track.TrackSingleFollowing(ctx, following)
	if err != nil {
		return fmt.Errorf("failed to track specific user %w", err)
	}

	stats := uc.osuAPI.GetTransportStats()
	defer uc.osuAPI.ResetStats()

	txErr := uc.txm.ReadWrite(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		if err = uc.log.Create(ctx, tx, &model.Log{
			Name:               fmt.Sprintf("Initial tracking for user %s", following.Username),
			Message:            model.LogMessageInitialTrack,
			Service:            "osu-dashboard-api",
			AppVersion:         "v1.0",
			Platform:           "Backend",
			Type:               model.LogTypeInitial,
			TrackedAt:          time.Now().UTC(),
			ElapsedTime:        time.Since(started),
			APIRequests:        stats.RequestCount,
			SuccessRatePercent: stats.SuccessRate,
			AvgResponseTime:    stats.AvgResponseTime,
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
