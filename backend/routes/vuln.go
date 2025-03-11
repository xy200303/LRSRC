package routes

import (
	"github.com/gin-gonic/gin"
	"xiaoyun/backend/controllers/vuln_controller"
	"xiaoyun/backend/middleware"
)

func SetupVulnRoutes(route *gin.Engine) {
	////不需要登录
	//{
	//	r := route.Group("/")
	//	r.GET(apiuri+"/getvulnabs", controllers.GetVulnAbstract)
	//	r.GET(apiuri+"/getvulntypes", controllers.GetVulnTypeList)
	//	r.GET(apiuri+"/getvulntype", controllers.GetVulnType)
	//	r.GET(apiuri+"/getvulnlist", controllers.GetVulnList)
	//	r.GET(apiuri+"/getvulndtl", controllers.GetVulnDetail)
	//}
	////需要登录
	//{
	//	r := route.Group("/")
	//	r.POST(apiuri+"/addvuln", controllers.AddVuln)
	//	r.POST(apiuri+"/editvuln", controllers.EditVuln)
	//	r.GET(apiuri+"/uservulnlist", controllers.GetUserVulnList)
	//}
	{
		r := route.Group(apiuri + "/vuln")
		r.GET("/buildVulnTypeTree", vuln_controller.BuildVulnTypeTree)

	}
	{
		r := route.Group(apiuri + "/vuln")
		r.Use(middleware.AuthMiddleware())
		r.POST("/submitVuln", vuln_controller.SubmitVuln)
	}
	////需要管理员
	{
		r := route.Group(apiuri + "/vuln")
		r.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
		r.POST("/listVulnType", vuln_controller.ListVulnType)
		r.POST("/createVulnType", vuln_controller.CreateVulnType)
		r.POST("/updateVulnType", vuln_controller.UpdateVulnType)
		r.POST("/deleteVulnType", vuln_controller.DeleteVulnType)

		//漏洞处理相关
		r.POST("/listVuln", vuln_controller.ListVuln)
		r.POST("/createVuln", vuln_controller.CreateVuln)
		r.POST("/updateVuln", vuln_controller.UpdateVuln)
		r.POST("/deleteVuln", vuln_controller.DeleteVuln)
		r.GET("/getVuln", vuln_controller.GetVuln)

		//审核漏洞
		r.POST("/auditVuln", vuln_controller.AuditVuln)

	}
}
