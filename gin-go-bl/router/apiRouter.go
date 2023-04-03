package router

import (
	v1 "gin-go-bl/api/v1"
	"gin-go-bl/middleware"
	"github.com/gin-gonic/gin"
)

func ApiRouter(r *gin.Engine) {
	r1 := r.Group("/v1")
	{
		//用户接口
		r1.POST("/user", v1.AddUser)
		r1.GET("/user/:num/:size", v1.GetUser)
		r1.PUT("/user/:id", middleware.AuthMiddleware(), v1.UpdateUser)
		r1.DELETE("/user/:id", middleware.AuthMiddleware(), v1.DeleteUser)

		//分类接口
		r1.POST("/cate", middleware.AuthMiddleware(), v1.AddCategory)
		r1.GET("/cate/:num/:size", v1.GetCategory)
		r1.PUT("/cate/:id", middleware.AuthMiddleware(), v1.UpdateCategory)
		r1.DELETE("/cate/:id", middleware.AuthMiddleware(), v1.DeleteCategory)

		//文章接口
		r1.POST("/article", middleware.AuthMiddleware(), v1.AddArticle)
		r1.GET("/articles/:num/:size", v1.GetAllArticle)
		r1.GET("/articles", v1.GetAllArticles)
		r1.GET("/article/:id", v1.GetArticle)
		r1.GET("/userarticle/:id/:num/:size", v1.GetUserArticle)
		r1.GET("/myarticle/:num/:size", middleware.AuthMiddleware(), v1.GetMeArticle)
		r1.GET("/catearticle/:id/:num/:size", v1.GetCaArticle)
		r1.PUT("/article/:id", middleware.AuthMiddleware(), v1.EditArticle)
		r1.DELETE("/article/:id", middleware.AuthMiddleware(), v1.DeleteArticle)

		//登录
		r1.POST("/login", v1.Login)

	}

}
