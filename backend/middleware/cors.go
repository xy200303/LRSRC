package middleware

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
	"xiaoyun/backend/config"
)

func CorsMiddleware() gin.HandlerFunc {
	// 配置CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = strings.Split(config.Config.Server.AllowOrigins, ",")
	corsConfig.AllowMethods = strings.Split(config.Config.Server.AllowMethods, ",")
	corsConfig.AllowHeaders = strings.Split(config.Config.Server.AllowHeaders, ",")
	// 添加调试输出
	fmt.Println("CORS Configuration:")
	fmt.Printf("AllowOrigins: %v\n", corsConfig.AllowOrigins)
	fmt.Printf("AllowMethods: %v\n", corsConfig.AllowMethods)
	fmt.Printf("AllowHeaders: %v\n", corsConfig.AllowHeaders)
	return cors.New(corsConfig)
}
