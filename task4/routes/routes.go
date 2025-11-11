package routes

import (
	"task4/handlers"
	"task4/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// 公开路由
	api := router.Group("/api")
	{
		// 认证路由
		auth := api.Group("/auth")
		{
			auth.POST("/register", handlers.Register)
			auth.POST("/login", handlers.Login)
		}

		// 文章路由（公开）
		posts := api.Group("/posts")
		{
			posts.GET("/", handlers.GetPosts)
			posts.GET("/:id", handlers.GetPost)
		}
	}

	// 需要认证的路由
	authenticated := api.Group("/")
	authenticated.Use(middleware.AuthMiddleware())
	{
		// 文章管理（需要认证）
		posts := authenticated.Group("/posts")
		{
			posts.POST("/", handlers.CreatePost)
			posts.PUT("/:id", middleware.AuthorizePost(), handlers.UpdatePost)
			posts.DELETE("/:id", middleware.AuthorizePost(), handlers.DeletePost)
		}

		// 评论管理
		comments := authenticated.Group("/comments")
		{
			comments.POST("/", handlers.CreateComment)
			comments.GET("/post/:postId", handlers.GetPostComments)
		}
	}
}
