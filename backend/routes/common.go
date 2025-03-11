package routes

import (
	"github.com/gin-gonic/gin"
	"xiaoyun/backend/controllers/common_controllers"
	"xiaoyun/backend/controllers/system_controllers"
	"xiaoyun/backend/middleware"
)

func SetupCommonRoutes(route *gin.Engine) {
	//不需要登录
	{
		r := route.Group("/")
		r.POST(apiuri+"/register", common.Register)
		r.POST(apiuri+"/login", common.Login)
		r.POST(apiuri+"/forgetPassword", common.ForgetPassword)
		//r.GET(apiuri+"/search", controllers.SearchVuln)
		//r.POST(apiuri+"/advsearch", controllers.SearchVulnAdv)
		//r.GET("/download/file", controllers.DownloadFile)
		r.GET(apiuri+"/getCaptcha", common.GetCaptcha)
		r.GET(apiuri+"/getDictType", system_controllers.GetDictType)
		r.GET(apiuri+"/getDictData", system_controllers.GetDictData)
		//r.POST(apiuri+"/forgetpassword", controllers.ForgetPassword)
		r.GET(apiuri+"/downloadFile", common.DownloadFile)
		r.GET(apiuri+"/listDictType", system_controllers.ListDictType)
		r.GET(apiuri+"/listDictData", system_controllers.ListDictData)
		r.GET(apiuri+"/getBaseSysConfigMap", system_controllers.GetBaseSysConfigMap)
	}
	//需要登录
	{
		r := route.Group("/")
		r.Use(middleware.AuthMiddleware())
		r.GET(apiuri+"/checkToken", common.CheckToken)
		r.GET(apiuri+"/getMyProfile", common.GetMyProfile)
		r.GET(apiuri+"/logout", common.Logout)
		r.POST(apiuri+"/uploadFile", common.UploadFile)
		r.POST(apiuri+"/changePassword", common.ChangePassword)
		r.POST(apiuri+"/updateProfile", common.UpdateProfile)

	}
	//需要管理员权限
	{

	}

}
