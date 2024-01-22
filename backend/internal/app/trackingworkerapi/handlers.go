package trackingworkerapi

import (
	"context"
	"time"
)

func (w *Worker) Start(ctx context.Context) func() error {
	finished := make(chan struct{}, 1)

	go func() {
		for {
			func() {
				//if w.tracker.GetLastTimeTracked

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
				w.lg.Infof("tracking finished")
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
