package userrepository

import (
	"context"
	"fmt"
	"playcount-monitor-backend/internal/repository/model"
)

const usersTableName = "users"

func (r *GormRepository) Create(ctx context.Context, user *model.User) error {
	err := r.db.WithContext(ctx).Table(usersTableName).Create(user).Error
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (r *GormRepository) Get(ctx context.Context, id string) (*model.User, error) {
	var user *model.User
	err := r.db.WithContext(ctx).Table(usersTableName).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get user with id %v: %w", id, err)
	}

	return user, nil
}

func (r *GormRepository) GetByName(ctx context.Context, name string) (*model.User, error) {
	var user *model.User
	err := r.db.WithContext(ctx).Table(usersTableName).Where("username = ?", name).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get user by name %s: %w", name, err)
	}

	return user, nil
}

func (r *GormRepository) Update(ctx context.Context, user *model.User) error {
	err := r.db.WithContext(ctx).Table(usersTableName).Save(user).Error
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

func (r *GormRepository) List(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	err := r.db.WithContext(ctx).Table(usersTableName).Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}

	return users, nil
}
