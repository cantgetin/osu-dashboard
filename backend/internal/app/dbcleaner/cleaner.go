package dbcleaner

import (
	"context"
	"time"
)

const cleanEveryNHours = 24 * time.Hour

func (w *Worker) Start(ctx context.Context) func() error {
	finished := make(chan struct{}, 1)

	lastTimeCleaned, err := w.cleaner.GetLastTimeCleaned(ctx)
	if err != nil {
		w.lg.Errorf("failed to get last time cleaned: %v", err)
		return func() error {
			<-finished
			return nil
		}
	}

	hoursSinceLastFetch := time.Since(*lastTimeCleaned).Hours()
	if hoursSinceLastFetch <= cleanEveryNHours.Hours() {
		waitDuration := time.Duration(cleanEveryNHours.Hours()-hoursSinceLastFetch) * time.Hour
		w.lg.Errorf("persisted last time cleaned:, waiting %v until next clean", waitDuration)
		time.Sleep(waitDuration)
	} else {
		w.lg.Infof("persisted last time cleaned: %v, no need to wait until refetch", *lastTimeCleaned)
	}

	go func() {
		for {
			func() {
				w.lg.Infof("cleaning worker started")

				loopCtx, cancel := context.WithTimeout(ctx, w.cfg.CleaningTimeout)
				defer cancel()

				if err := w.cleaner.Clean(loopCtx); err != nil {
					w.lg.Errorf("encountered error while cleaning: %v", err)
					return
				}

				err = w.cleaner.CreateCleanRecord(ctx)
				if err != nil {
					w.lg.Errorf("failed to create clen record: %v", err)
				}

				w.lg.Infof("cleaned successfully")
			}()

			select {
			case <-ctx.Done():
				finished <- struct{}{}
				w.lg.Infof("cleaning finished")
				return
			case <-time.After(w.cfg.CleaningInterval):
			}
		}
	}()

	return func() error {
		<-finished
		return nil
	}
}
