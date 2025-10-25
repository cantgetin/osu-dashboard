package osuapi

import (
	"context"
	"golang.org/x/sync/errgroup"
	"strconv"
	"sync"
)

func (s *Service) GetUserWithMapsets(ctx context.Context, userID string) (*User, []*MapsetExtended, error) {
	eg, egCtx := errgroup.WithContext(ctx)

	var (
		user            *User
		mapsets         []*Mapset
		mapsetsExtended []*MapsetExtended
		mu              = sync.Mutex{}
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

	// now we wanna get comments count for each mapset hehe
	// can be a lot of mapsets here so limit parallel calls
	semaphore := make(chan struct{}, s.cfg.TrackingMaxParallelMapsetCalls)

	for _, mapset := range mapsets {
		eg.Go(func() (errG error) {
			semaphore <- struct{}{}
			defer func() {
				<-semaphore
			}()

			commentCount, errG := s.GetMapsetCommentsCount(egCtx, strconv.Itoa(mapset.Id))
			if errG != nil {
				return errG
			}

			mu.Lock()
			defer mu.Unlock()
			mapsetsExtended = append(mapsetsExtended, &MapsetExtended{
				Mapset:        mapset,
				CommentsCount: commentCount,
			})

			return nil
		})
	}

	err = eg.Wait()
	if err != nil {
		return nil, nil, err
	}

	return user, mapsetsExtended, nil
}
