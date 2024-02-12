package trackingworkerapi

import (
	"context"
	"time"
)

func (w *Worker) Start(ctx context.Context) func() error {
	finished := make(chan struct{}, 1)

	lastTimeTracked, err := w.tracker.GetLastTimeTracked(ctx)
	if err != nil {
		w.lg.Errorf("failed to get last time tracked: %v", err)
		return func() error {
			<-finished
			return nil
		}
	}

	hoursSinceLastFetch := time.Now().Sub(*lastTimeTracked).Hours()
	if hoursSinceLastFetch <= 24 {
		waitDuration := time.Duration(24-hoursSinceLastFetch) * time.Hour
		w.lg.Errorf("persisted last time tracked:, waiting %v until next fetch", waitDuration)
		time.Sleep(waitDuration)
	} else {
		w.lg.Infof("persisted last time tracked: %v, no need to wait until refetch", *lastTimeTracked)
	}

	go func() {
		for {
			func() {
				w.lg.Infof("tracking worker started")

				loopCtx, cancel := context.WithTimeout(ctx, w.cfg.TrackingTimeout)
				defer cancel()

				if err := w.tracker.Track(loopCtx); err != nil {
					w.lg.Errorf("encountered error while tracking: %v", err)
					return
				}
				w.lg.Infof("tracked successfully")
			}()

			select {
			case <-ctx.Done():
				finished <- struct{}{}
				createCtx, cancelCreateCtx := context.WithTimeout(context.Background(), time.Duration(time.Second*5))
				err = w.tracker.CreateTrackRecord(createCtx)
				if err != nil {
					w.lg.Errorf("failed to create track record: %v", err)
				}
				w.lg.Infof("tracking finished")
				cancelCreateCtx()
				return
			case <-time.After(w.cfg.TrackingInterval):
			}
		}
	}()

	return func() error {
		<-finished
		return nil
	}
}
