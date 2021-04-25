package routes

import (
	"ginBlog/api/v1"
	"ginBlog/config"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(config.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		//user模块的路由接口
		router.POST("user", v1.AddUser)
		router.GET("user", v1.GetUsers)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DeleteUser)

		//分类模块的路由接口
		router.POST("category", v1.AddCategory)
		router.GET("category", v1.GetCategory)
		router.PUT("category/:id", v1.EditCategory)
		router.DELETE("category/:id", v1.DeleteCategory)

		//文章模块的路由接口
		router.POST("article", v1.AddArticle)
		router.GET("article", v1.GetArticles)
		router.GET("article/info/:id", v1.GetArticleInfo)
		router.GET("article/list/:cid", v1.GetCateArt)
		router.PUT("article/:id", v1.EditArticle)
		router.DELETE("article/:id", v1.DeleteArticle)

	}
	return r
}
