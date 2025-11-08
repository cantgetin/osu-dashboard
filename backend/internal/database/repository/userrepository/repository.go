package userrepository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"osu-dashboard/internal/database/repository/model"
	"osu-dashboard/internal/database/txmanager"
	"strings"
)

const usersTableName = "users"

func (r *GormRepository) Create(ctx context.Context, tx txmanager.Tx, user *model.User) error {
	err := tx.DB().WithContext(ctx).Table(usersTableName).Create(user).Error
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (r *GormRepository) Get(ctx context.Context, tx txmanager.Tx, id int) (*model.User, error) {
	var user *model.User
	err := tx.DB().WithContext(ctx).Table(usersTableName).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get user with id %v: %w", id, err)
	}

	return user, nil
}

func (r *GormRepository) Exists(ctx context.Context, tx txmanager.Tx, id int) (bool, error) {
	var count int64
	err := tx.DB().WithContext(ctx).Table(usersTableName).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, fmt.Errorf("failed to check if user with id %v exists: %w", id, err)
	}

	return count > 0, nil
}

func (r *GormRepository) GetByName(ctx context.Context, tx txmanager.Tx, name string) (*model.User, error) {
	var user *model.User
	err := tx.DB().WithContext(ctx).Table(usersTableName).Where("username = ?", name).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get user by name %s: %w", name, err)
	}

	return user, nil
}

func (r *GormRepository) Update(ctx context.Context, tx txmanager.Tx, user *model.User) error {
	err := tx.DB().WithContext(ctx).Table(usersTableName).Save(user).Error
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

func (r *GormRepository) List(ctx context.Context, tx txmanager.Tx) ([]*model.User, error) {
	var users []*model.User
	err := tx.DB().WithContext(ctx).Table(usersTableName).Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}

	return users, nil
}

func (r *GormRepository) ListUsersWithFilterSortLimitOffset(
	ctx context.Context,
	tx txmanager.Tx,
	filter model.UserFilter,
	sort model.UserSort,
	limit int,
	offset int,
) ([]*model.User, int, error) {
	var users []*model.User
	var count int64

	query, values := buildListByFilterQuery(filter)
	filterGormExpr := gorm.Expr(query, values...)
	if query == "" {
		filterGormExpr = gorm.Expr("1 = 1")
	}

	err := tx.DB().WithContext(ctx).
		Table(usersTableName).
		Where(filterGormExpr).
		Count(&count).Error
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count users: %w", err)
	}

	order := buildOrderBySortQuery(sort)
	if len(strings.TrimSpace(order)) == 0 {
		order = "created_at DESC"
	}

	err = tx.DB().WithContext(ctx).
		Table(usersTableName).
		Order(order).
		Where(filterGormExpr).
		Limit(limit).
		Offset(offset).
		Find(&users).Error
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list users: %w", err)
	}

	return users, int(count), nil
}

func (r *GormRepository) TotalCount(ctx context.Context, tx txmanager.Tx) (int, error) {
	var count int64
	err := tx.DB().WithContext(ctx).Table(usersTableName).Count(&count).Error
	if err != nil {
		return 0, fmt.Errorf("failed to count users: %w", err)
	}

	return int(count), nil
}

func buildListByFilterQuery(filter model.UserFilter) (string, []interface{}) {
	if len(filter) == 0 {
		return "", nil
	}

	for field, value := range filter {
		switch field {
		case model.UserNameField:
			username, ok := value.(string)
			if !ok {
				return "", nil
			}
			return "username ILIKE ?", []interface{}{"%" + username + "%"}

		default:
			return "", nil
		}
	}

	return "", nil
}

func buildOrderBySortQuery(sort model.UserSort) string {
	// thats evil stuff but idc since users table is tiny, wont affect query performance much
	return fmt.Sprintf(
		"(user_stats->(SELECT MAX(k) FROM jsonb_object_keys(user_stats) AS k)->>'%s')::INT %s",
		string(sort.Field),
		string(sort.Direction),
	)
}
