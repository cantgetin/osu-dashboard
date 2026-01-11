package mapsetrepository

import (
	"context"
	"fmt"
	"osu-dashboard/internal/database/repository/model"
	"osu-dashboard/internal/database/txmanager"
	gosort "sort"
	"strings"

	"gorm.io/gorm"
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
	err := tx.DB().WithContext(ctx).Table(mapsetsTableName).Omit("genre", "language").Save(mapset).Error
	if err != nil {
		return fmt.Errorf("failed to update mapset: %w", err)
	}

	return nil
}

func (r *GormRepository) UpdateFull(ctx context.Context, tx txmanager.Tx, mapset *model.Mapset) error {
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

func (r *GormRepository) TotalCount(ctx context.Context, tx txmanager.Tx) (int, error) {
	var count int64
	err := tx.DB().WithContext(ctx).Table(mapsetsTableName).Count(&count).Error
	if err != nil {
		return 0, fmt.Errorf("failed to count mapsets: %w", err)
	}

	return int(count), nil
}

func (r *GormRepository) UpdateGenreLanguage(
	ctx context.Context, tx txmanager.Tx, id int, newGenre string, newLanguage string,
) error {
	// select mapset and switch its genre and language inside a transaction
	mapset, err := r.Get(ctx, tx, id)
	if err != nil {
		return err
	}

	mapset.Genre = newGenre
	mapset.Language = newLanguage

	return r.UpdateFull(ctx, tx, mapset)
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

func (r *GormRepository) ListWithFilterAndLimit(
	ctx context.Context,
	tx txmanager.Tx,
	filter model.MapsetFilter,
	limit int,
) ([]*model.Mapset, error) {
	var mapsets []*model.Mapset

	query, values := buildListByFilterQuery(filter)
	filterGormExpr := gorm.Expr(query, values...)
	if query == "" {
		filterGormExpr = gorm.Expr("1 = 1")
	}

	err := tx.DB().WithContext(ctx).
		Table(mapsetsTableName).
		Where(filterGormExpr).
		Limit(limit).
		Find(&mapsets).Error
	if err != nil {
		return nil, fmt.Errorf("failed to list mapsets: %w", err)
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
) ([]*model.Mapset, int, error) {
	var mapsets []*model.Mapset
	var count int64

	query, values := buildListByFilterQuery(filter)
	filterGormExpr := gorm.Expr(query, values...)
	if query == "" {
		filterGormExpr = gorm.Expr("1 = 1")
	}

	order := buildOrderBySortQuery(sort)
	if strings.TrimSpace(order) == "" {
		order = "created_at DESC"
	}

	err := tx.DB().WithContext(ctx).
		Table(mapsetsTableName).
		Where(filterGormExpr).
		Count(&count).Error

	if err != nil {
		return nil, 0, fmt.Errorf("failed to count mapsets: %w", err)
	}

	err = tx.DB().WithContext(ctx).
		Table(mapsetsTableName).
		Order(order).
		Where(filterGormExpr).
		Limit(limit).
		Offset(offset).
		Find(&mapsets).Error

	if err != nil {
		return nil, 0, fmt.Errorf("failed to list mapsets: %w", err)
	}

	return mapsets, int(count), nil
}

func (r *GormRepository) ListForUserWithFilterSortLimitOffset(
	ctx context.Context,
	tx txmanager.Tx,
	userID int,
	filter model.MapsetFilter,
	sort model.MapsetSort,
	limit int,
	offset int,
) ([]*model.Mapset, int, error) {
	var mapsets []*model.Mapset
	var count int64

	query, values := buildListByFilterQuery(filter)
	filterGormExpr := gorm.Expr(query, values...)
	if query == "" {
		filterGormExpr = gorm.Expr("1 = 1")
	}

	order := buildOrderBySortQuery(sort)
	if strings.TrimSpace(order) == "" {
		order = "created_at DESC"
	}

	err := tx.DB().WithContext(ctx).
		Table(mapsetsTableName).
		Where(filterGormExpr).
		Where("user_id = ?", userID).
		Count(&count).Error

	if err != nil {
		return nil, 0, fmt.Errorf("failed to count mapsets: %w", err)
	}

	err = tx.DB().WithContext(ctx).
		Table(mapsetsTableName).
		Order(order).
		Where(filterGormExpr).
		Where("user_id = ?", userID).
		Limit(limit).
		Offset(offset).
		Find(&mapsets).Error

	if err != nil {
		return nil, 0, fmt.Errorf("failed to list mapsets: %w", err)
	}

	return mapsets, int(count), nil
}

func buildListByFilterQuery(filter model.MapsetFilter) (query string, values []any) {
	var queryBuilder strings.Builder
	values = make([]any, 0, len(filter))
	keys := make([]string, 0, len(filter))
	for column := range filter {
		keys = append(keys, string(column))
	}

	gosort.Strings(keys)

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
