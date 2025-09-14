package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"myproject/web/service"
	"myproject/web/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strconv"
	"myproject/web/model"
	"gorm.io/gorm/logger"
    "log"
	"os"
	"time"
	"myproject/web/middleware"
	"myproject/web/dto"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	
	userService := service.NewUserService(repository.NewUserRepo(db))
	postService := service.NewPostService(repository.NewPostRepo(db))
	commentService := service.NewCommentService(repository.NewCommentRepo(db), repository.NewPostRepo(db))

	
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/getUser/:userId", func(c *gin.Context) {
		uidstr := c.Params.ByName("userId")
		uid, _ := strconv.ParseInt(uidstr, 10, 64)
		nuid := uint(uid)
		user, _:= userService.GetUser(c.Request.Context(), nuid)
		c.JSON(http.StatusOK, user)
	})

	r.POST("/register", func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		tuser,err := userService.Register(c.Request.Context(), user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, tuser)
	})

	r.POST("/login", func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		token,err := userService.Login(c.Request.Context(), user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.String(http.StatusOK, token)
	})

	r.Use(middleware.AuthMiddleware())
	{
		r.POST("/publishPost", func(c *gin.Context) {

			var postReq dto.PublishPostRequest
			if err := c.ShouldBindJSON(&postReq); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			userId,_ := c.Get("userID")
			post := model.Post{ID: postReq.PostId, Content: postReq.Content, Title: postReq.Title, UserId: userId.(uint)}
			result,err := postService.SavePost(c.Request.Context(), post)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, result)
		})

		r.POST("/listPosts", func(c *gin.Context) {

			var queryReq dto.Pagination[dto.QueryPostListRequest]
			if err := c.ShouldBindJSON(&queryReq); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			result,err := postService.QueryList(c.Request.Context(), &queryReq)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, result)
		})

		r.DELETE("/removePost/:postId", func(c *gin.Context) {

			postIdStr := c.Params.ByName("postId")
			tpostId, _ := strconv.ParseInt(postIdStr, 10, 64)
			postId := uint(tpostId)
			if postId == 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "文章找不到"})
				return
			}
			result,err := postService.DeleteById(c.Request.Context(), postId)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"success": result})
		})

		r.POST("/addComment", func(c *gin.Context) {

			var addComment dto.AddComment
			if err := c.ShouldBindJSON(&addComment); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if addComment.PostId == 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "文章找不到"})
				return
			}
			userId,_ := c.Get("userID")
			result,err := commentService.AddComment(c.Request.Context(), &addComment, userId.(uint))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, result)
		})

		r.POST("/listComment", func(c *gin.Context) {
			var queryReq dto.Pagination[uint]
			if err := c.ShouldBindJSON(&queryReq); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if queryReq.QueryParameter == 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "文章找不到"})
				return
			}
			result,err := commentService.QueryList(c.Request.Context(), &queryReq)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, result)
		})
	}
	

	return r


}

func main() {
	// logger, _ := zap.NewProduction()
    // defer logger.Sync()

    // r := gin.New()

    // // Gin 中间件输出 zap 日志
    // r.Use(func(c *gin.Context) {
    //     logger.Info("request",
    //         zap.String("path", c.Request.URL.Path),
    //         zap.String("method", c.Request.Method),
    //     )
    //     c.Next()
    // })

	newLogger := logger.New(
        log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
        logger.Config{
            SlowThreshold: time.Second,   // 慢查询阈值
            LogLevel:      logger.Info,   // 日志级别 (Silent, Error, Warn, Info)
            Colorful:      true,          // 彩色打印
        },
    )
	var db,err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: newLogger,
	})


	db.Exec("DROP TABLE posts")
	db.Exec("DROP TABLE comments")
	db.Exec("DROP TABLE users")
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Post{})
	db.AutoMigrate(&model.Comment{})

	if err != nil {
		panic("failed to connect database.")
	}
	router := setupRouter(db)

	router.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

