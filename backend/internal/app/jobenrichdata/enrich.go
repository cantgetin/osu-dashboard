package enricher

import (
	"context"
	"time"
)

const enrichEveryNHours = 168 * time.Hour

func (w *Worker) Start(ctx context.Context) func() error {
	finished := make(chan struct{}, 1)

	lastTimeenriched, err := w.enricher.GetLastTimeEnriched(ctx)
	if err != nil {
		w.lg.Errorf("failed to get last time enrich: %v", err)
		return func() error {
			<-finished
			return nil
		}
	}

	hoursSinceLastFetch := time.Since(*lastTimeenriched).Hours()
	if hoursSinceLastFetch <= enrichEveryNHours.Hours() {
		waitDuration := time.Duration(enrichEveryNHours.Hours()-hoursSinceLastFetch) * time.Hour
		w.lg.Errorf("persisted last time enriched:, waiting %v until next enrich", waitDuration)
		time.Sleep(waitDuration)
	} else {
		w.lg.Infof("persisted last time enriched: %v, no need to wait until refetch", *lastTimeenriched)
	}

	go func() {
		for {
			func() {
				w.lg.Infof("enriching worker started")

				loopCtx, cancel := context.WithTimeout(ctx, w.cfg.EnrichingTimeout)
				defer cancel()

				if err := w.enricher.Enrich(loopCtx); err != nil {
					w.lg.Errorf("encountered error while enriching: %v", err)
					return
				}

				err = w.enricher.CreateEnrichRecord(ctx)
				if err != nil {
					w.lg.Errorf("failed to create enrich record: %v", err)
				}

				w.lg.Infof("enriched successfully")
			}()

			select {
			case <-ctx.Done():
				finished <- struct{}{}
				w.lg.Infof("enriching finished")
				return
			case <-time.After(w.cfg.EnrichingInterval):
			}
		}
	}()

	return func() error {
		<-finished
		return nil
	}
}
