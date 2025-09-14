package service

import (
	"context"
    "myproject/web/model"
	"myproject/web/dto"
    "myproject/web/repository"
)

type PostService struct {
	repo repository.PostRepo
}

func NewPostService(repo repository.PostRepo) *PostService {
	return &PostService{repo: repo}
}

func (p *PostService) SavePost(ctx context.Context, post model.Post) (*model.Post, error) {
	var id uint
	var err error
	if post.ID != 0 {
		id, err = p.repo.Update(ctx, &post)
	} else {
		id, err = p.repo.Create(ctx, &post)
	}
	if err != nil {
		return nil, err
	}
	return p.repo.FindById(ctx, id)
}

func (p *PostService) QueryList(ctx context.Context, req *dto.Pagination[dto.QueryPostListRequest]) (*dto.PageResult[dto.PostDTO], error) {

	return p.repo.QueryList(ctx, req)
}

func (p *PostService) DeleteById(ctx context.Context, id uint) (bool, error){
	return p.repo.DeleteById(ctx, id)
}