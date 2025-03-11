package routes

import (
	"github.com/gin-gonic/gin"
	"xiaoyun/backend/controllers/article_controllers"
	"xiaoyun/backend/middleware"
)

func SetupArticleRoutes(route *gin.Engine) {
	//需要登录
	{
		r := route.Group(apiuri + "/article")
		r.Use(middleware.AuthMiddleware())
		//r.GET(apiuri+"/getSysBaseInfo", system_controllers.GetSysBaseInfo)
		r.POST("/sendArticleComment", article_controllers.SendArticleComment)
		r.POST("/listArticleComment", article_controllers.ListArticleComment)
		//删除自己的评论
		r.POST("/deleteMyArticleComment", article_controllers.DeleteMyArticleComment)
		//获取首页数据
		r.GET("/getArticleGroup", article_controllers.GetArticleGroup)
		r.GET("/getArticle", article_controllers.GetArticle)
		r.POST("/listHomeArticle", article_controllers.ListHomeArticle)
		r.POST("/listHomeAllArticleGroup", article_controllers.ListHomeAllArticleGroup)
	}
	//需要管理员
	{
		r := route.Group(apiuri + "/article")
		r.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
		r.POST("/listArticleGroup", article_controllers.ListArticleGroup)
		r.POST("/deleteArticleGroup", article_controllers.DeleteArticleGroup)
		r.POST("/updateArticleGroup", article_controllers.UpdateArticleGroup)
		r.POST("/createArticleGroup", article_controllers.CreateArticleGroup)
		r.POST("/ListAllArticleGroup", article_controllers.ListAllArticleGroup)

		//发表文章
		r.POST("/createArticle", article_controllers.CreateArticle)
		r.POST("/updateArticle", article_controllers.UpdateArticle)
		r.POST("/listArticle", article_controllers.ListArticle)
		r.POST("/deleteArticle", article_controllers.DeleteArticle)
		//文章评论
		r.POST("/deleteArticleComment", article_controllers.DeleteArticleComment)

	}

}
