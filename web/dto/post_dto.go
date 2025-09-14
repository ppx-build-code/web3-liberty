package dto


type PublishPostRequest struct {
	PostId uint
	Content string `json:"content" binding:"required"`
	Title string
}

type QueryPostListRequest struct {
	PostId [] uint
	KeyWord string
}

type PostDTO struct {
	PostId uint
	Content string
	Title string
}

// PageResult 通用分页返回结构
type PageResult[T any] struct {
    List      []T   `json:"list"`
    Total     int64 `json:"total"`
    PageNum      int   `json:"pageNum"`
    PageSize  int   `json:"pageSize"`
    TotalPage int   `json:"totalPage"`
}

// Pagination 分页请求参数
type Pagination[T any] struct {
    PageNum     int `form:"pageNum" json:"pageNum"`
    PageSize int `form:"pageSize" json:"pageSize"`
	QueryParameter T
}