package repository

import (
	"myproject/web/model"
	"myproject/web/dto"
	"context"
)

type CommentRepo interface {

	QueryList(ctx context.Context, req *dto.Pagination[uint]) (*dto.PageResult[model.Comment], error)

	Add(ctx context.Context, comment *model.Comment) (*model.Comment, error)
}