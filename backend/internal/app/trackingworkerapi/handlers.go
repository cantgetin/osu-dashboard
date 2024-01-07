package trackingworkerapi

import (
	"context"
	"time"
)

func (w *Worker) Start(ctx context.Context) func() error {
	finished := make(chan struct{}, 1)

	go func() {
		w.lg.Infof("following worker started")

		for {
			func() {
				w.lg.Infof("following worker started")

				loopCtx, cancel := context.WithTimeout(ctx, w.cfg.TrackingTimeout)
				defer cancel()

				if err := w.tracker.Track(loopCtx); err != nil {
					w.lg.Errorf("encountered error while following: %v", err)
					return
				}
				w.lg.Infof("tracked successfully")
			}()

			select {
			case <-ctx.Done():
				finished <- struct{}{}
				w.lg.Infof("following finished")
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
