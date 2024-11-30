package repository

import (
	"context"
	"simple-dashboard-server/model"

	"gorm.io/gorm"
)

type UserRepo interface {
	FindByID(ctx context.Context, id string) (model.User, error)
	FindByEmail(ctx context.Context, email string) (model.User, error)
	UpdateByID(ctx context.Context, id string, updates map[string]interface{}) error
	Create(ctx context.Context, user model.User) error
}

type userRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		DB: db,
	}
}

func (r *userRepo) FindByID(ctx context.Context, id string) (model.User, error) {
	var user model.User
	err := r.DB.WithContext(ctx).
		Where("id = ?", id).
		First(&user).Error
	return user, err
}

func (r *userRepo) FindByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User
	err := r.DB.WithContext(ctx).
		Where("email = ?", email).
		First(&user).Error
	return user, err
}

func (r *userRepo) UpdateByID(ctx context.Context, id string, updates map[string]interface{}) error {
	err := r.DB.WithContext(ctx).
		Model(&model.User{}).
		Where("id = ?", id).
		Updates(&updates).Error
	return err
}

func (r *userRepo) Create(ctx context.Context, user model.User) error {
	err := r.DB.WithContext(ctx).
		Create(&user).Error
	return err
}
