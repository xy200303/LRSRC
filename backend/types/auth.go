package types

import (
	"xiaoyun/backend/types/user_types"
)

// Auth 验证鉴权对象
type Auth struct {
	User    *user_types.User
	RoleId  uint8
	IsAdmin bool
}

// LoginReq 登录请求接口
type LoginReq struct {
	LoginType string `json:"login_type"`
	Username  string `json:"username" binding:"required,min=3,max=20"`
	Password  string `json:"password" binding:"required"`
}

// RegisterReq 注册接口
type RegisterReq struct {
	RegisterType string `json:"register_type"`
	Nickname     string `json:"nickname" binding:"required"`
	Username     string `json:"username" binding:"required,username,min=3,max=20"`
	Password     string `json:"password" binding:"required,min=6"`
	Email        string `json:"email" binding:"required,email"`
	Phone        string `json:"phone" binding:"required,phone"`
	Code         string `json:"code" binding:"required,min=4"`
}

// ForgetPwdReq 忘记密码接口
type ForgetPwdReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Code     string `json:"code" binding:"required,min=4"`
}

// LoginResp 登录相应实体
type LoginResp struct {
	Username string `json:"username"`
	RoleId   uint8  `json:"role_id"`
	Token    string `json:"token"`
}

// ChangePasswordReq 修改密码请求
type ChangePasswordReq struct {
	OldPassword string `json:"old_password" binding:"required,min=6"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}
