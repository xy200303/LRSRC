package routes

import (
	"github.com/gin-gonic/gin"
	"xiaoyun/backend/controllers/system_controllers"
	"xiaoyun/backend/middleware"
)

func SetupSystemRoutes(route *gin.Engine) {
	//不需要登录
	{
		r := route.Group("/")
		r.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
		//r.GET(apiuri+"/getSysBaseInfo", system_controllers.GetSysBaseInfo)
	}
	//需要管理员
	{
		r := route.Group("/")
		r.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
		r.GET(apiuri+"/getSysConfigMap", system_controllers.GetSysConfigMap)
		r.GET(apiuri+"/getSysStatus", system_controllers.GetSystemStatus)
		r.POST(apiuri+"/updateSysConfigMap", system_controllers.UpdateSysConfigMap)
		r.POST(apiuri+"/deleteDictType", system_controllers.DeleteDictType)
		r.POST(apiuri+"/createDictType", system_controllers.CreateDictType)
		r.POST(apiuri+"/updateDictType", system_controllers.UpdateDictType)

		r.POST(apiuri+"/deleteDictData", system_controllers.DeleteDictData)
		r.POST(apiuri+"/createDictData", system_controllers.CreateDictData)
		r.POST(apiuri+"/updateDictData", system_controllers.UpdateDictData)

	}

}
