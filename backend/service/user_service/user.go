package user_service

import (
	"fmt"
	"xiaoyun/backend/models"
	"xiaoyun/backend/types/user_types"
	"xiaoyun/backend/utils"
)

// ListUser 用户列表
func ListUser(req *user_types.UserQuery) ([]*user_types.User, int64, error) {
	users, total, err := models.ListUser(req)
	if err != nil {
		return users, total, err
	}
	return users, total, nil
}

// SetAdminUser 设置管理员用户
func SetAdminUser(req *user_types.UserAdminReq) error {
	err := models.SetAdminUser(&user_types.User{
		Username: req.Username,
		IsAdmin:  req.IsAdmin,
	})
	return err
}

// CreateUser 创建用户
func CreateUser(req *user_types.UserReq) error {
	//验证手机号是否存在
	if IsUserExist(req.Username) {
		return fmt.Errorf("用户名已经存在")
	}
	if IsUserExist(req.Email) {
		return fmt.Errorf("邮箱已经存在")
	}
	if IsUserExist(req.Phone) {
		return fmt.Errorf("手机号已经存在")
	}
	err := models.CreateUser(&user_types.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  utils.GenPasswordHash("123456"),
		Phone:     req.Phone,
		Gender:    req.Gender,
		Name:      req.Name,
		IDCard:    req.IDCard,
		Birthdate: req.Birthdate,
		Address:   req.Address,
		Status:    req.Status,
		Nickname:  req.Nickname,
	})
	if err != nil {
		return err
	}
	return nil
}

// 删除用户
func DeleteUser(username []string) error {
	err := models.DeleteUser(username)
	if err != nil {
		return err
	}
	return nil
}

// UpdateUser 更新用户
func UpdateUser(req *user_types.UserReq) error {
	err := models.UpdateUser(&user_types.User{
		Username:  req.Username,
		Email:     req.Email,
		Phone:     req.Phone,
		Gender:    req.Gender,
		Name:      req.Name,
		IDCard:    req.IDCard,
		Birthdate: req.Birthdate,
		Address:   req.Address,
		Status:    req.Status,
		Nickname:  req.Nickname,
	})
	return err
}

// IsUserExist 检查用户是否存在
func IsUserExist(param string) bool {
	_, err := models.FindUser(param)
	if err != nil {
		return false
	}
	return true
}
