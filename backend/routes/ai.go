package routes

import (
	"github.com/gin-gonic/gin"
	"xiaoyun/backend/controllers/ai_controller"
	"xiaoyun/backend/middleware"
)

func SetupAiRoutes(route *gin.Engine) {

	{
		r := route.Group(apiuri + "/ai")
		r.POST("/chat/completions", ai_controller.CommonChatStream)
	}
	//需要登录
	{
		r := route.Group(apiuri + "/ai")
		r.Use(middleware.AuthMiddleware())
		r.POST("/summaryContent", ai_controller.SummaryContent)
		r.POST("/chat", ai_controller.CommonChat)
		//r.POST("/chat/completions", ai_controller.CommonChatStream)
	}
	//需要管理员权限
	{

	}

}
