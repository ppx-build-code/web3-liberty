
package main

import (
	"context"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"fmt"
)

type User struct {
	// gorm.Model
	ID uint
	Username string
	Passowrd string
	PostCount uint32
	Posts [] Post
}

type Post struct {
	// gorm.Model
	ID uint
	UserId uint
	Title string
	Comments [] Comment
	CommentStatusDesc string
}

type Comment struct {
	// gorm.Model
	ID uint
	PostId uint
	Content string
}

type Stat struct {
	Pid uint
	Cnum uint
}

func (post *Post) AfterCreate(tx *gorm.DB) (err error) {
	// ctx := tx.Statement.Context
	// gorm.G[Post](tx).Where("id = ?", post.ID).Update("post_count", gorm.Expr("post_count + ?", 1))
	fmt.Printf("触发钩子函数+1,%v,\n",post.ID)
	result := tx.Model(User{}).Where("id = ?", post.UserId).Update("post_count", gorm.Expr("post_count + ? ", 1))
	return result.Error
}

func (comment *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var count int64
	result := tx.Model(Comment{}).Where("post_id = ?", comment.PostId).Count(&count)
	fmt.Printf("触发钩子函数-移除评论,%v,剩余数量:%v,\n",comment, count)
	if count <= 0 {
		tx.Save(&Post{ID: comment.PostId, CommentStatusDesc: "无评论"})
		var post Post
		tx.First(&post, comment.PostId)
		fmt.Printf("触发钩子函数-更新状态后，文章数据:%v,\n",post)
	}
	
	return result.Error
}

func main() {
	db,err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	}

	db.Exec("DROP TABLE posts")
	db.Exec("DROP TABLE comments")
	db.Exec("DROP TABLE users")

	user := User{
		Username:"张三",
		Passowrd:"123456", 
		Posts: []Post{
			{
				Title:"first post",
				Comments: []Comment{
					{Content: "hello1"},
					{Content: "hello2"},
					{Content: "hello3"},
				},
			},
			{
				Title:"second post",
				Comments: []Comment{
					{Content: "hello1"},
					{Content: "hello2"},
					{Content: "hello3"},
					{Content: "hello4"},
				},
			},
		}}

	posts:= []*Post{
		{
			UserId: 1,
			Title:"first post",
			Comments: []Comment{
				{Content: "hello1"},
				{Content: "hello2"},
				{Content: "hello3"},
			},
		},
		{
			UserId: 1,
			Title:"second post",
			Comments: []Comment{
				{Content: "hello1"},
				{Content: "hello2"},
				{Content: "hello3"},
				{Content: "hello4"},
			},
		},
	}
	// post := Post{UserId: 1, Title:"first"}
	// comment := Comment{PostId: 1, Title:"first"}
	ctx := context.Background()
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Comment{})

	gorm.G[User](db).Create(ctx,&user)
	db.Create(&posts)


	// gorm.G[Post](db).Create(ctx)
	// gorm.G[Comment](db).Create(ctx)
	// gorm.G[Post](db).Create(ctx,&post)

	var tuser User
	result1, _ := gorm.G[Post](db).Where("id = ?", 1).Find(ctx)
	fmt.Printf("data: %v\n", result1)
	db.Model(&User{}).Where("id = 1").Preload("Posts").Preload("Posts.Comments").Find(&tuser)
	fmt.Printf("data: %v\n", tuser)

	var count int64
	db.Table("posts").Count(&count)
	fmt.Printf("Posts count:%v,\n", count)

	var res Stat
	db.Table("posts a").
	Select("a.id as pid, count(b.id) as cnum").
	Joins("left join Comments b on a.id = b.post_id").
	Group("a.id").
	Order("cnum desc").
	Limit(1).
	Scan(&res)
	fmt.Printf("data: %v\n", res)

	// var postCount uint32
	var luser User
	db.First(&luser, 1)
	// db.Table("Users").Select("post_count").Scan(&postCount)
	fmt.Printf("postCount: %v\n", luser)

	var firstComments [] Comment
	firstComments, _ = gorm.G[Comment](db).Where("post_id = ?", 1).Find(ctx)
	// result := db.Find(&firstComments)
	fmt.Printf("firstComment: %v, 行数: %v\n", firstComments)


	
	
	db.Delete(&firstComments)

	// cols, _ := db.Migrator().ColumnTypes("posts")
	// for _, col := range cols {
    // 	fmt.Println("\ncolumn:", col.Name())
	// }

}