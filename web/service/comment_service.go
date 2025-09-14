package service

import (
	"myproject/web/repository"
	"myproject/web/model"
	"myproject/web/dto"
	"errors"
	"context"

)

type CommentService struct {
	commentRepo repository.CommentRepo
	postRepo repository.PostRepo
}

func NewCommentService(repo repository.CommentRepo, postRepo repository.PostRepo) *CommentService {
	return &CommentService{
		commentRepo: repo,
		postRepo: postRepo,
	}
}

func (c *CommentService) AddComment(ctx context.Context, comment *dto.AddComment, userId uint) (*model.Comment, error) {
	post, _ := c.postRepo.FindById(ctx, comment.PostId)
	if post == nil {
		return nil, errors.New("not found post")
	}
	data := model.Comment{
		Content: comment.Content,
		PostId: comment.PostId,
		UserId: userId,
	}
	return c.commentRepo.Add(ctx, &data)
}

func (c *CommentService) QueryList(ctx context.Context, req *dto.Pagination[uint]) (*dto.PageResult[model.Comment], error){
	return c.commentRepo.QueryList(ctx, req)

}