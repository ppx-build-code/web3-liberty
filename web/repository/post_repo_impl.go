package repository

import (
	"context"
	"myproject/web/model"
	"gorm.io/gorm"
	"myproject/web/utils"
	"myproject/web/dto"
	// "fmt"
)


type postRepo struct {
	db *gorm.DB
}

func NewPostRepo(db *gorm.DB) PostRepo {
	return &postRepo{db: db}
}

func (r *postRepo) Create(ctx context.Context, post *model.Post) (uint, error) {
	result := r.db.WithContext(ctx).Create(&post)
	return post.ID, result.Error
}

func (r *postRepo) Update(ctx context.Context, post *model.Post) (uint, error) {
	result := r.db.WithContext(ctx).Updates(&post)
	return post.ID, result.Error
}

func (r *postRepo) FindById(ctx context.Context, id uint) (*model.Post, error) {
	var post model.Post
	if err := r.db.WithContext(ctx).First(&post, id).Error; err != nil {
		return &post, err
	}
	return &post, nil
}

func (r *postRepo) QueryList(ctx context.Context, req *dto.Pagination[dto.QueryPostListRequest]) (*dto.PageResult[dto.PostDTO], error) {
	var posts [] model.Post
	query := r.db.WithContext(ctx).Model(&model.Post{})
	if (len(req.QueryParameter.PostId) >0) {
		query.Where("id in ?", req.QueryParameter.PostId)
	}
	if (req.QueryParameter.KeyWord != "") {
		query.Where("title like ?", "%" + req.QueryParameter.KeyWord + "%")
	}
	var count int64
	query.Count(&count)
	if err := query.Scopes(utils.Paginate(req)).Find(&posts).Error; err != nil {
		return nil, err
	}

	

	dataList := make([] dto.PostDTO, 0)
	for _, post := range posts {
		v := dto.PostDTO{
			PostId: post.ID,
			Content: post.Content,
			Title: post.Title,
			
		}
		dataList = append(dataList, v)
	}
	res := dto.PageResult[dto.PostDTO]{
		List: dataList,
		PageNum: req.PageNum,
		PageSize: req.PageSize,
		Total: count,
	}
	return &res, nil
}

func (r *postRepo) DeleteById(ctx context.Context, id uint) (bool, error) {
	result := r.db.WithContext(ctx).Delete(&model.Post{}, id)
	if result.RowsAffected <= 0 {
		return false, result.Error
	}
	return true, result.Error
}
