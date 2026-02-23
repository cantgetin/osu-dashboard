package jobrecordusecase

import (
	"context"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"osu-dashboard/internal/config"
	"osu-dashboard/internal/database/model"
	jobrepository "osu-dashboard/internal/database/repository/job"
	"osu-dashboard/internal/database/txmanager"
	"time"
)

type JobType string

const (
	JobTypeTrackUsers JobType = "track_users"
	JobTypeCleanStats JobType = "clean_stats"
	JobTypeCleanUsers JobType = "clean_users"
	JobTypeEnrichData JobType = "enrich_data"
)

type UseCase struct {
	jobType JobType
	cfg     *config.Config
	lg      *log.Logger
	txm     txmanager.TxManager
	jobs    JobStore
}

type JobStore interface {
	Create(ctx context.Context, tx txmanager.Tx, job *model.Job) error
	GetLastForType(ctx context.Context, tx txmanager.Tx, jobType string) (*model.Job, error)
}

func New(
	cfg *config.Config,
	lg *log.Logger,
	jobType JobType,
	txm txmanager.TxManager,
	jobs JobStore,
) *UseCase {
	return &UseCase{
		jobType: jobType,
		cfg:     cfg,
		lg:      lg,
		txm:     txm,
		jobs:    jobs,
	}
}

func (uc *UseCase) Record(ctx context.Context, startTime, endTime time.Time, err error) error {
	hasError := err != nil
	errorText := ""
	if err != nil {
		errorText = err.Error()
	}
	typedJob := &model.Job{
		Name:      string(uc.jobType),
		StartedAt: startTime,
		EndedAt:   endTime,
		Error:     hasError,
		ErrorText: errorText,
	}

	return uc.txm.ReadWrite(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		if err = uc.jobs.Create(ctx, tx, typedJob); err != nil {
			return fmt.Errorf("failed to record job, err: %v", err)
		}
		return nil
	})
}

func (uc *UseCase) GetLastRecordStartTime(ctx context.Context) (time.Time, error) {
	lastJob := new(model.Job)
	txErr := uc.txm.ReadOnly(ctx, func(ctx context.Context, tx txmanager.Tx) (err error) {
		lastJob, err = uc.jobs.GetLastForType(ctx, tx, string(uc.jobType))
		if err != nil {
			if errors.Is(err, jobrepository.RecordNotFound) {
				lastJob.StartedAt = time.Now()
				return nil
			}

			return fmt.Errorf("failed to get last record for type %v")
		}

		return nil
	})
	return lastJob.StartedAt, txErr
}
