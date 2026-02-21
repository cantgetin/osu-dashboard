package jobrepository

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"osu-dashboard/internal/config"
	"osu-dashboard/internal/database/model"
	"osu-dashboard/internal/database/txmanager"

	log "github.com/sirupsen/logrus"
)

var RecordNotFound = fmt.Errorf("record not found")

const jobsTableName = "jobs"

type GormRepository struct {
	lg  *log.Logger
	cfg *config.Config
}

type Interface interface {
	Create(ctx context.Context, tx txmanager.Tx, job *model.Job) error
	GetLastForType(ctx context.Context, tx txmanager.Tx, jobType string) (*model.Job, error)
}

func New(cfg *config.Config, lg *log.Logger) *GormRepository {
	return &GormRepository{
		lg:  lg,
		cfg: cfg,
	}
}

func (r *GormRepository) Create(ctx context.Context, tx txmanager.Tx, job *model.Job) error {
	// todo: fix this auto increment crap
	var maxID *int
	err := tx.DB().WithContext(ctx).Table(jobsTableName).Select("max(id)").Row().Scan(&maxID)
	if err != nil {
		return fmt.Errorf("failed to create job: %w", err)
	}
	job.ID = *maxID + 1

	err = tx.DB().WithContext(ctx).Table(jobsTableName).Create(job).Error
	if err != nil {
		return fmt.Errorf("failed to create job: %w", err)
	}

	return nil
}

func (r *GormRepository) GetLastForType(ctx context.Context, tx txmanager.Tx, jobType string) (*model.Job, error) {
	var job model.Job
	err := tx.DB().
		WithContext(ctx).
		Table(jobsTableName).
		Where("name = ?", jobType).
		Order("started_at desc").
		Limit(1).
		Find(&job).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, RecordNotFound
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get last job: %w", err)
	}

	return &job, nil
}
