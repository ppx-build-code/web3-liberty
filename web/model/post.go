package model

import "time"

type Post struct {
	ID uint	`gorm:primaryKey`
	Title string
	Content string
	UserId uint
	CreateAt time.Time
	UpdateAt time.Time
	Comments [] Comment
}