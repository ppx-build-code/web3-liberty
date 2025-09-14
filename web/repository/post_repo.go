package repository

import (
	"myproject/web/model"
	"context"
	"myproject/web/dto"
)

type PostRepo interface {

	Create(ctx context.Context, post *model.Post) (uint, error)

	Update(ctx context.Context, post *model.Post) (uint, error)

	FindById(ctx context.Context, id uint) (*model.Post, error)
	
	DeleteById(ctx context.Context, id uint) (bool, error)

	QueryList(ctx context.Context, req *dto.Pagination[dto.QueryPostListRequest]) (*dto.PageResult[dto.PostDTO], error)
}