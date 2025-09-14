package model

import "time"

type Comment struct {
	ID uint
	Content string
	UserId uint
	PostId uint
	CreateAt time.Time
}