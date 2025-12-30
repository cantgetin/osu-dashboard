package enricherusecase

import (
	"context"
	"fmt"
	"osu-dashboard/internal/database/repository/model"
	"osu-dashboard/internal/database/txmanager"
	"strconv"
	"time"
)

func (uc *UseCase) Enrich(ctx context.Context) error {
	startTime := time.Now()

	// getting all users IDs
	var follows []*model.Following
	if err := uc.txm.ReadOnly(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		var err error
		if follows, err = uc.following.List(ctx, tx); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	// log
	uc.lg.Infof("got following IDs from db, %v total", len(follows))
	if len(follows) == 0 {
		return fmt.Errorf("no following users present in db")
	}

	for _, f := range follows {
		if err := uc.enrichUserMapsets(ctx, f.ID); err != nil {
			return fmt.Errorf("failed to enrich user mapsets: %w", err)
		}
	}

	defer uc.osuApi.ResetStats()
	return uc.CreateEnrichAndLogRecords(ctx, startTime)
}

func (uc *UseCase) enrichUserMapsets(ctx context.Context, userID int) error {
	// get user mapsets
	var userMapsets []*model.Mapset
	if err := uc.txm.ReadOnly(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		var err error
		if userMapsets, err = uc.mapset.ListForUser(ctx, tx, userID); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	// fetch extended info for each mapset in case if user added/switched genre, language of mapset
	for _, mapset := range userMapsets {
		langGenreInfo, err := uc.osuApi.GetMapsetExtended(ctx, strconv.Itoa(mapset.ID))
		if err != nil {
			return fmt.Errorf("failed to get mapset extended info from api, mapset id: %v, err: %w", mapset.ID, err)
		}

		newGenre := langGenreInfo.Genre.Name
		newLanguage := langGenreInfo.Language.Name

		// update mapset genre, language inside readwrite transaction
		txErr := uc.txm.ReadWrite(ctx, func(ctx context.Context, tx txmanager.Tx) error {
			if err = uc.mapset.UpdateGenreLanguage(ctx, tx, mapset.ID, newGenre, newLanguage); err != nil {
				return fmt.Errorf("failed to update genre language for mapset: %v, err: %w", mapset.ID, err)
			}
			return nil
		})
		if txErr != nil {
			return fmt.Errorf("failed to update genre lang, transaction failed: err: %w", err)
		}
	}
	return nil
}

func (uc *UseCase) GetLastTimeEnriched(ctx context.Context) (*time.Time, error) {
	t := time.Time{}
	if err := uc.txm.ReadOnly(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		enrich, err := uc.enriches.GetLastEnrich(ctx, tx)
		if err != nil {
			return err
		}

		t = enrich.EnrichedAt

		return nil
	}); err != nil {
		return nil, err
	}

	return &t, nil
}

func (uc *UseCase) CreateEnrichRecord(ctx context.Context) error {
	enrich := &model.Enrich{
		EnrichedAt: time.Now().UTC(),
	}

	if err := uc.txm.ReadWrite(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		err := uc.enriches.Create(ctx, tx, enrich)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) CreateEnrichAndLogRecords(ctx context.Context, startTime time.Time) error {
	var (
		elapsed       = time.Since(startTime)
		reqs          = uc.osuApi.GetOutgoingRequestCount()
		respTime      = uc.osuApi.AverageResponseTime()
		avgReqsPerMin = float64(reqs) / elapsed.Minutes()
		successRate   = uc.osuApi.SuccessRate()
	)

	uc.lg.Infof("Sent %v requests to api in %v minutes", reqs, elapsed.Minutes())
	uc.lg.Infof("Average requests per minute: %f", avgReqsPerMin)

	if err := uc.CreateEnrichRecord(ctx); err != nil {
		return fmt.Errorf("failed to create enrich record: %w", err)
	}

	txErr := uc.txm.ReadWrite(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		if err := uc.log.Create(ctx, tx, &model.Log{
			Name:               "Daily mapset info enrichment",
			Message:            model.LogMessageDailyEnrich,
			Service:            "enricher",
			AppVersion:         "v1.0",
			Platform:           "Backend",
			Type:               model.LogTypeRegular,
			APIRequests:        reqs,
			SuccessRatePercent: successRate,
			TrackedAt:          time.Now().UTC(),
			AvgResponseTime:    respTime,
			ElapsedTime:        elapsed,
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
