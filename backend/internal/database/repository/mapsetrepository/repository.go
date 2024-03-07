package mapsetrepository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
	"sort"
	"strings"
)

const mapsetsTableName = "mapsets"

func (r *GormRepository) Create(ctx context.Context, tx txmanager.Tx, mapset *model.Mapset) error {
	err := tx.DB().WithContext(ctx).Table(mapsetsTableName).Create(mapset).Error
	if err != nil {
		return fmt.Errorf("failed to create mapset: %w", err)
	}

	return nil
}

func (r *GormRepository) Get(ctx context.Context, tx txmanager.Tx, id int) (*model.Mapset, error) {
	var mapset *model.Mapset
	err := tx.DB().WithContext(ctx).Table(mapsetsTableName).Where("id = ?", id).First(&mapset).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get mapset with id %v: %w", id, err)
	}

	return mapset, nil
}

func (r *GormRepository) Update(ctx context.Context, tx txmanager.Tx, mapset *model.Mapset) error {
	err := tx.DB().WithContext(ctx).Table(mapsetsTableName).Save(mapset).Error
	if err != nil {
		return fmt.Errorf("failed to update mapset: %w", err)
	}

	return nil
}

func (r *GormRepository) Exists(ctx context.Context, tx txmanager.Tx, id int) (bool, error) {
	var count int64
	err := tx.DB().WithContext(ctx).Table(mapsetsTableName).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, fmt.Errorf("failed to check if mapset with id %v exists: %w", id, err)
	}

	return count > 0, nil
}

func (r *GormRepository) List(ctx context.Context, tx txmanager.Tx) ([]*model.Mapset, error) {
	var mapsets []*model.Mapset
	err := tx.DB().WithContext(ctx).Table(mapsetsTableName).Find(&mapsets).Error
	if err != nil {
		return nil, fmt.Errorf("failed to list mapsets: %w", err)
	}

	return mapsets, nil
}

func (r *GormRepository) ListForUser(ctx context.Context, tx txmanager.Tx, userID int) ([]*model.Mapset, error) {
	var mapsets []*model.Mapset
	err := tx.DB().WithContext(ctx).Table(mapsetsTableName).Where("user_id = ?", userID).Find(&mapsets).Error
	if err != nil {
		return nil, fmt.Errorf("failed to list mapsets for user %v: %w", userID, err)
	}

	return mapsets, nil
}

func (r *GormRepository) ListStatusesForUser(
	ctx context.Context,
	tx txmanager.Tx,
	userID int,
) ([]string, error) {
	var statuses []string
	err := tx.DB().WithContext(ctx).Table(mapsetsTableName).
		Where("user_id = ?", userID).
		Pluck("status", &statuses).Error

	if err != nil {
		return nil, fmt.Errorf("failed to list statuses for user %v: %w", userID, err)
	}

	return statuses, nil
}

func (r *GormRepository) ListForUserWithLimitOffset(
	ctx context.Context,
	tx txmanager.Tx,
	userID int,
	limit int,
	offset int,
) ([]*model.Mapset, error) {
	var mapsets []*model.Mapset
	err := tx.DB().WithContext(ctx).Table(mapsetsTableName).
		Where("user_id = ?", userID).
		Order("last_playcount DESC").
		Limit(limit).
		Offset(offset).
		Find(&mapsets).Error

	if err != nil {
		return nil, fmt.Errorf("failed to list mapsets for user %v: %w", userID, err)
	}

	return mapsets, nil
}

func (r *GormRepository) ListWithFilterSortLimitOffset(
	ctx context.Context,
	tx txmanager.Tx,
	filter model.MapsetFilter,
	sort model.MapsetSort,
	limit int,
	offset int,
) ([]*model.Mapset, error) {
	var mapsets []*model.Mapset

	query, values := buildListByFilterQuery(filter)
	filterGormExpr := gorm.Expr(query, values...)
	if query == "" {
		filterGormExpr = gorm.Expr("1 = 1")
	}

	order := buildOrderBySortQuery(sort)
	if len(strings.TrimSpace(order)) == 0 {
		order = "created_at DESC"
	}

	err := tx.DB().WithContext(ctx).
		Table(mapsetsTableName).
		Order(order).
		Where(filterGormExpr).
		Limit(limit).
		Offset(offset).
		Find(&mapsets).Error

	if err != nil {
		return nil, fmt.Errorf("failed to list mapsets: %w", err)
	}

	return mapsets, nil
}

func buildListByFilterQuery(filter model.MapsetFilter) (string, []interface{}) {
	var queryBuilder strings.Builder
	values := make([]interface{}, 0, len(filter))
	keys := make([]string, 0, len(filter))
	for column := range filter {
		keys = append(keys, string(column))
	}

	sort.Strings(keys)

	for i, column := range keys {
		if column == string(model.MapsetArtistOrTitleOrTagsFields) {
			columns := []model.MapsetFilterField{model.MapsetArtistField, model.MapsetTitleField, model.MapsetTagsField}

			// search
			queryBuilder.WriteString("( ")
			for i, c := range columns {
				if i == 0 {
					queryBuilder.WriteString(string(c) + " ILIKE ?")
				} else {
					queryBuilder.WriteString(" OR " + string(c) + " ILIKE ?")
				}
				values = append(values, "%"+filter[model.MapsetFilterField(column)].(string)+"%")
			}
			queryBuilder.WriteString(" )")

		} else {
			queryBuilder.WriteString(column + " = ?")
			values = append(values, filter[model.MapsetFilterField(column)])
		}
		if i < len(filter)-1 {
			queryBuilder.WriteString(" AND ")
		}
	}

	return queryBuilder.String(), values
}

func buildOrderBySortQuery(sort model.MapsetSort) string {
	return fmt.Sprintf("%s %s", string(sort.Field), string(sort.Direction))
}
