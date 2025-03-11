package user_types

import (
	"time"
	"xiaoyun/backend/types/base_types"
)

// User 用户表
type User struct {
	Username  string    `gorm:"primaryKey;type:varchar(40)" json:"username"`
	Password  string    `json:"password"`
	Email     string    `gorm:"unique;index;type:varchar(32)" json:"email"`
	Nickname  string    `gorm:"type:varchar(40)" json:"nickname"`
	Gender    string    `gorm:"default:0" json:"gender"`
	Integral  int32     `gorm:"default:0" json:"integral"` // 积分值
	Phone     string    `gorm:"unique;index;type:varchar(20)" json:"phone"`
	Level     uint8     `gorm:"default:0" json:"level"`
	RoleID    uint8     `gorm:"default:0" json:"role_id"` // 角色ID
	Birthdate time.Time `gorm:"type:date" json:"birthdate"`
	Address   string    `gorm:"type:varchar(200)" json:"address"`
	Name      string    `gorm:"type:varchar(20)" json:"name"` // 用户姓名
	IDCard    string    `gorm:"type:char(18)" json:"id_card"` // 身份证号
	Status    string    `gorm:"default:1" json:"status"`
	Avatar    string    `json:"avatar"`
	IsAdmin   bool      `gorm:"default:0" json:"is_admin"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserReq 用户表，用于创建和修改用户
type UserReq struct {
	Username  string    `json:"username" binding:"required,username,min=3,max=20"`
	Password  string    `json:"password"`
	Email     string    `json:"email" binding:"required,email"`
	Nickname  string    `json:"nickname" binding:"required"`
	Gender    string    `json:"gender"`
	Integral  int32     `json:"integral"` // 积分值
	Phone     string    `json:"phone" binding:"required"`
	Level     uint8     `json:"level"`
	RoleID    uint8     `json:"role_id"` // 角色ID
	Birthdate time.Time `json:"birthdate"`
	Address   string    `json:"address"`
	Name      string    `json:"name"`    // 用户姓名
	IDCard    string    `json:"id_card"` // 身份证号
	Status    string    `json:"status"`
	Avatar    string    `json:"avatar"`
}

// UserQuery 查询用户参数表
type UserQuery struct {
	base_types.BasePage
	Username string `json:"username"`
	Email    string `json:"email" `
	Nickname string `json:"nickname" `
	Phone    string `json:"phone"`
}

// UserAdminReq 设置管理员参数
type UserAdminReq struct {
	IsAdmin  bool   `json:"is_admin" binding:"omitempty"`
	Username string `json:"username" binding:"required,username,min=3,max=20"`
}
