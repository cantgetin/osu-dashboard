package job

import (
	"context"
	log "github.com/sirupsen/logrus"
	"time"
)

type (
	PeriodicJob struct {
		Name     string
		Log      *log.Logger
		Period   time.Duration
		Timeout  time.Duration
		Recorder Recorder
		Executor Executor
	}

	Recorder interface {
		Record(ctx context.Context, startTime, endTime time.Time, err error) error
		GetLastRecordStartTime(ctx context.Context) (time.Time, error)
	}

	Executor interface {
		Execute(ctx context.Context) error
	}
)

func NewPeriodic(lg *log.Logger, name string, period, timeout time.Duration, r Recorder, e Executor) *PeriodicJob {
	return &PeriodicJob{
		Log:      lg,
		Name:     name,
		Period:   period,
		Timeout:  timeout,
		Recorder: r,
		Executor: e,
	}
}

func (pj *PeriodicJob) Start(ctx context.Context) {
	go func() {
		for {
			func() {
				pj.Log.Infof("job %v started", pj.Name)

				lastRecord, err := pj.Recorder.GetLastRecordStartTime(ctx)
				if err != nil {
					pj.Log.Errorf("failed to get %v job last record, err: %v", pj.Name, err)
					return
				}

				timeSinceLastRecord := time.Since(lastRecord)
				if timeSinceLastRecord <= pj.Period {
					waitDuration := pj.Period - timeSinceLastRecord
					pj.Log.Errorf("persisted last record for %v job:, waiting %v until next execute", pj.Name, waitDuration)
					time.Sleep(waitDuration)
				} else {
					pj.Log.Infof("persisted last record for %v job: %v, no need to wait until refetch", pj.Name, lastRecord)
				}

				loopCtx, cancel := context.WithTimeout(ctx, pj.Timeout)
				defer cancel()

				startTime := time.Now()
				err = pj.Executor.Execute(loopCtx)
				if err != nil {
					pj.Log.Errorf("encountered error while executing %v job, err: %v", pj.Name, err)
				}

				endTime := time.Now()

				err = pj.Recorder.Record(ctx, startTime, endTime, err)
				if err != nil {
					pj.Log.Errorf("failed to record %v job result, err: %v", pj.Name, err)
				}

				pj.Log.Infof("successfully executed and recorded %v job", pj.Name)
			}()

			select {
			case <-ctx.Done():
				pj.Log.Infof("received context done for %v job", pj.Name)
				return
			case <-time.After(pj.Period):
				pj.Log.Infof("waited for defined period of %v for %v job", pj.Period, pj.Name)
			}
		}
	}()

	return
}
