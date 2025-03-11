package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
	"xiaoyun/backend/config"
	"xiaoyun/backend/middleware"
	"xiaoyun/backend/models"
)

var (
	route  *gin.Engine
	apiuri = "/api/v1"
)

func init() {
	switch config.Config.Server.Mode {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	case "debug":
		gin.SetMode(gin.DebugMode)
	}
	route = gin.Default()
	fmt.Println("Welcome to XiaoYun", config.Version)
	fmt.Println("Server running on " + config.Config.Server.Host + ":" + strconv.FormatInt(config.Config.Server.Port, 10))
	//初始化数据库和JWT参数
	err := models.RedisSetJwtSecret("jwtToken")
	if err != nil {
		log.Fatalf("jwt_secret err: %v", err)
	}
	jwtToken, _ := models.GetRedisValue("jwt_secret")
	log.Printf("JWT_TOKEN:", jwtToken)

	//设置文件上传目录
	if err = os.MkdirAll(config.Config.Server.UploadDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create uploads directory: %v", err)
	}
}

func SetupRouter() {
	//注册通用接口
	SetupCommonRoutes(route)
	//注册系统接口
	SetupSystemRoutes(route)
	//注册用户接口
	SetupUserRoutes(route)
	//文章接口注册
	SetupArticleRoutes(route)
	//ai接口注册
	SetupAiRoutes(route)
	//注册漏洞管理路由
	SetupVulnRoutes(route)
	//SetupVulnRoutes(route)
	//SetupUserRoutes(route)
}

// 前后端分离的路由
func Api() {
	route.Use(middleware.CorsMiddleware())
	SetupRouter()
	err := route.Run(config.Config.Server.Host + ":" + strconv.FormatInt(config.Config.Server.Port, 10))
	if err != nil {
		return
	}
}

// 前后端不分离的路由
func All() {
	webPath := config.Config.Server.WebPath
	webStaticUrl := config.Config.Server.StaticUrl
	route.Use(middleware.CorsMiddleware())
	SetupRouter()
	// 前端静态文件
	log.Printf(webPath)
	route.Static(webStaticUrl, webPath+webStaticUrl)

	// 设置前端模板文件的路由
	route.LoadHTMLGlob(webPath + "/*.html")

	// 定义路由
	route.GET("/", func(c *gin.Context) {
		c.File(webPath + "/index.html")
	})

	// 通配符路由
	route.NoRoute(func(c *gin.Context) {
		c.File(webPath + "/index.html")
	})

	err := route.Run(config.Config.Server.Host + ":" + strconv.FormatInt(config.Config.Server.Port, 10))
	if err != nil {
		fmt.Println("启动服务端失败", err)
	}
}
