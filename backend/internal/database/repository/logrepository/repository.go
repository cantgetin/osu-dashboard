package logrepository

import (
	"context"
	"fmt"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
)

const logTableName = "log"

func (r *GormRepository) Create(ctx context.Context, tx txmanager.Tx, log *model.Log) error {
	log.ID = 0
	err := tx.DB().WithContext(ctx).Table(logTableName).Create(log).Error
	if err != nil {
		return fmt.Errorf("failed to create log: %w", err)
	}
	return nil
}

func (r *GormRepository) List(ctx context.Context, tx txmanager.Tx) ([]*model.Log, error) {
	var logs []*model.Log
	err := tx.DB().WithContext(ctx).Table(logTableName).Find(&logs).Error
	if err != nil {
		return nil, fmt.Errorf("failed to list logs: %w", err)
	}

	return logs, nil
}

func (r *GormRepository) ListLogsWithLimitOffset(
	ctx context.Context,
	tx txmanager.Tx,
	limit,
	offset int,
) (logs []*model.Log, count int, err error) {
	var cnt int64

	err = tx.DB().WithContext(ctx).
		Table(logTableName).
		Count(&cnt).Error
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count logs: %w", err)
	}

	err = tx.DB().WithContext(ctx).
		Table(logTableName).
		Order("tracked_at desc").
		Limit(limit).
		Offset(offset).
		Find(&logs).Error
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list logs: %w", err)
	}

	return logs, int(cnt), nil
}
