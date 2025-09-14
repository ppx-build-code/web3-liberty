package repository

import (
	"context"
	"myproject/web/model"
	"gorm.io/gorm"
	"fmt"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db : db}
}

func (r *userRepo) FindById(ctx context.Context, id uint) (*model.User, error) {
	var user model.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) FindByName(ctx context.Context, name string) (*model.User, error) {
	var user model.User
	if err := r.db.WithContext(ctx).First(&user, "username = ?", name).Error; err != nil {
		return nil, err
	}
	fmt.Printf("FindByName:%v", &user)
	return &user, nil
}

func (r *userRepo) Create(ctx context.Context, user *model.User) (uint,error) {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}
func (r *userRepo) Update(ctx context.Context, user *model.User) error {
	return nil
}
func (r *userRepo) Delete(ctx context.Context, user *model.User) error {
	return nil
}