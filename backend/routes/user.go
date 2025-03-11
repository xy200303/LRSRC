package routes

import (
	"github.com/gin-gonic/gin"
	"xiaoyun/backend/controllers/user_controllers"
	"xiaoyun/backend/middleware"
)

func SetupUserRoutes(route *gin.Engine) {
	////无需登录
	//{
	//	r := route.Group("/")
	//	r.GET(apiuri+"/usertop", controllers.GetUserTop10)
	//}
	//
	////需要登录
	//{
	//	r := route.Group("/")
	//	r.GET(apiuri+"/userinfo", controllers.GetUserInfo)
	//	r.POST(apiuri+"/updateavatar", controllers.UpdateAvatar)
	//	r.POST(apiuri+"/updateuserinfo", controllers.UpdateUserInfo)
	//	r.POST(apiuri+"/updatepassword", controllers.UpdateUserPassword)
	//}
	//
	//需要管理员权限
	{
		r := route.Group(apiuri + "/user")
		r.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
		r.POST("/list", user_controllers.ListUser)
		r.POST("/setUserAdmin", user_controllers.SetAdminUser)
		r.POST("/create", user_controllers.CreateUser)
		r.POST("/update", user_controllers.UpdateUser)
		r.POST("/delete", user_controllers.DeleteUser)
		//r.POST(apiuri+"/multidelusers", controllers.MultiDeleteUsers)
		//r.GET(apiuri+"/getusers", controllers.GetUsers)
		//r.POST(apiuri+"/updateuser", controllers.UpdateUser)
	}
}
