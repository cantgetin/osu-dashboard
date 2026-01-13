package osuapi

import (
	"context"
	osuapimodels "osu-dashboard/internal/service/osuapi/models"
	"strconv"
	"sync"

	"golang.org/x/sync/errgroup"
)

func (s *Service) GetUserWithMapsets(ctx context.Context, userID string) (*osuapimodels.User, []*osuapimodels.MapsetExtended, error) {
	eg, egCtx := errgroup.WithContext(ctx)

	var (
		user    *osuapimodels.User
		mapsets []*osuapimodels.Mapset
	)

	eg.Go(func() (errG error) {
		if user, errG = s.GetUser(egCtx, userID); errG != nil {
			return errG
		}
		return nil
	})

	eg.Go(func() (errG error) {
		if mapsets, errG = s.GetUserMapsets(egCtx, userID); errG != nil {
			return errG
		}
		return nil
	})

	err := eg.Wait()
	if err != nil {
		return nil, nil, err
	}

	mapsetsExtended, err := s.getMapsetsWithComments(ctx, mapsets)
	if err != nil {
		return nil, nil, err
	}

	return user, mapsetsExtended, nil
}

func (s *Service) getMapsetsWithComments(
	ctx context.Context, mapsets []*osuapimodels.Mapset,
) ([]*osuapimodels.MapsetExtended, error) {
	eg, egCtx := errgroup.WithContext(ctx)
	eg.SetLimit(s.cfg.TrackingMaxParallelMapsetCalls)

	var (
		mu              = sync.Mutex{}
		mapsetsExtended []*osuapimodels.MapsetExtended
	)

	for _, mapset := range mapsets {
		eg.Go(func() (errG error) {
			commentCount, errG := s.GetMapsetCommentsCount(egCtx, strconv.Itoa(mapset.Id))
			if errG != nil {
				return errG
			}

			mu.Lock()
			defer mu.Unlock()
			mapsetsExtended = append(mapsetsExtended, &osuapimodels.MapsetExtended{
				Mapset:        mapset,
				CommentsCount: commentCount,
			})

			return nil
		})
	}

	err := eg.Wait()
	if err != nil {
		return nil, err
	}

	return mapsetsExtended, nil
}
