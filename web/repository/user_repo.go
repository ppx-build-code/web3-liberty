package repository

import (
	"context"
	"myproject/web/model"
 )

 type UserRepo interface {
	FindById(ctx context.Context, id uint) (*model.User, error)
	FindByName(ctx context.Context, name string) (*model.User, error)
	Create(ctx context.Context, user *model.User)  (uint,error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, user *model.User) error
 }