package dto

type AddComment struct {
	PostId uint
	Content string `json:"content" binding:"required"`
}