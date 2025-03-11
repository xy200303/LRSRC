package middleware

import "C"
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"xiaoyun/backend/models"
	"xiaoyun/backend/types"
	"xiaoyun/backend/types/user_types"
	"xiaoyun/backend/utils"
	"xiaoyun/backend/utils/resp"
)

// GetCurrentUser 获取当前用户的结构体数据
func GetCurrentUser(c *gin.Context) (*types.Auth, error) {
	var auth types.Auth
	currentUser, ok := c.Get("user")
	if !ok {
		return nil, fmt.Errorf("用户信息不存在")
	}
	auth.User = currentUser.(*user_types.User)
	auth.RoleId = auth.User.RoleID
	auth.IsAdmin = auth.User.IsAdmin
	return &auth, nil
}

// AuthMiddleware 是一个自定义的身份验证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//method := c.Request.Method
		//path := c.FullPath()
		token := utils.GetToken(c)
		if token == "" {
			resp.Resp(c, 401, "无效的token", nil)
			return
		}
		//解码Token
		jwtSecret, err := models.RedisGetJwtSecret()
		if err != nil {
			resp.Resp(c, 401, "服务器密钥错误", nil)
			return
		}
		decodeJWT, err := utils.DecodeJWTToken(token, jwtSecret)
		if err != nil {
			resp.Resp(c, 401, "无效的token", nil)
			log.Println(err)
			return
		}
		fmt.Println(decodeJWT)
		//判断是否在redis中缓存
		redisToken, err := models.RedisGetUserToken(decodeJWT["username"].(string))
		if err != nil {
			resp.Resp(c, 401, "无效的token", nil)
			log.Println(err)
			return
		}
		if redisToken != token {
			resp.Resp(c, 401, "登录已经过期，请重新登录", nil)
		}
		//开始解析信息
		user, err := models.FindUserByUsername(decodeJWT["username"].(string))
		if err != nil {
			resp.Resp(c, 401, "无效的用户", nil)
			return
		}
		// 将用户信息存储在 gin.Context 中
		c.Set("user", user)
		// 继续处理请求
		c.Next()
	}
}

// AdminMiddleware Admin 是一个自定义的管理员身份验证中间件
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth, err := GetCurrentUser(c)
		if err != nil {
			resp.Resp(c, 401, err.Error(), nil)
			return
		}
		if !auth.IsAdmin {
			resp.Resp(c, 401, "非管理员用户禁止操作", nil)
			return
		}
		c.Next()
	}
}
