package repository

import (
	"myproject/web/model"
	"myproject/web/dto"
	"myproject/web/utils"
	"context"
	"gorm.io/gorm"
)

type commentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) CommentRepo {
	return &commentRepo{db: db}
}

func (c *commentRepo) QueryList(ctx context.Context, req *dto.Pagination[uint]) (*dto.PageResult[model.Comment], error) {
	var comments []model.Comment
	var total int64
	c.db.WithContext(ctx).Where("post_id = ?", req.QueryParameter).Count(&total)
	if err := c.db.WithContext(ctx).Scopes(utils.Paginate(req)).Where("post_id = ?", req.QueryParameter).Find(&comments).Error; err != nil {
		return nil, err
	}

	result := dto.PageResult[model.Comment]{
		List: comments,
		PageNum: req.PageNum,
		PageSize: req.PageSize,
		Total: total,
	}
	return &result, nil

}

func (c *commentRepo) Add(ctx context.Context, comment *model.Comment) (*model.Comment, error) {
	c.db.WithContext(ctx).Create(comment)
	return comment, nil
}