package utils

import (
	"myproject/web/dto"
	"gorm.io/gorm"
)

func Paginate[T any](r *dto.Pagination[T]) func(db *gorm.DB) *gorm.DB {
  return func (db *gorm.DB) *gorm.DB {
    page := r.PageNum
    if page <= 0 {
      page = 1
    }

    pageSize := r.PageSize
    switch {
    case pageSize > 100:
      pageSize = 100
    case pageSize <= 0:
      pageSize = 10
    }

    offset := (page - 1) * pageSize
    return db.Offset(offset).Limit(pageSize)
  }
}