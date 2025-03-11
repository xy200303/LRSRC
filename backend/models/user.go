package models

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"xiaoyun/backend/types/user_types"
)

// FindUserByUsername 寻找用户
func FindUserByUsername(username string) (*user_types.User, error) {
	var user user_types.User
	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// FindUser 寻找用户
func FindUser(param string) (*user_types.User, error) {
	var user user_types.User
	result := db.Where("username = ?", param).
		Or("email = ?", param).
		Or("phone = ?", param).
		First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("用户名或密码验证错误")
		}
		return nil, result.Error // 其他错误
	}
	return &user, nil
}

// CreateUser 创建用户
func CreateUser(req *user_types.User) error {
	// 用户名不存在，创建新用户
	result := db.Create(req)
	if result.Error != nil {
		return result.Error // 返回错误
	}
	return nil // 返回新创建的用户
}

// UpdateUser 更新数据
func UpdateUser(req *user_types.User) error {
	// 列出所有需要更新的字段
	fieldsToUpdate := []string{
		"email", "nickname", "gender", "phone",
		"birthdate", "address", "name", "id_card", "status", "avatar",
	}
	result := db.Model(&user_types.User{}).Where("username = ?", req.Username).
		Select(fieldsToUpdate).
		Updates(req)
	if result.Error != nil {
		return result.Error // 处理错误
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("未找到匹配的记录") // 没有更新任何记录
	}
	return nil // 更新成功
}

// 更新密码
func UpdatePassword(req *user_types.User) error {
	result := db.Model(&user_types.User{}).Where("username = ?", req.Username).
		Select("password").
		Updates(req)
	if result.Error != nil {
		return result.Error // 处理错误
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("未找到匹配的记录") // 没有更新任何记录
	}
	return nil // 更新成功
}

// DeleteUser 删除用户
func DeleteUser(usernames []string) error {
	// 检查 usernames 是否包含 "admin"
	for _, username := range usernames {
		if username == "admin" {
			return fmt.Errorf("无法删除管理员账户: admin") // 返回错误
		}
	}
	result := db.Where("username IN (?)", usernames).Delete(&user_types.User{})
	if result.Error != nil {
		return result.Error // 处理错误
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("未找到匹配的记录") // 没有删除任何记录
	}
	return nil // 删除成功
}

// ListUser 查询字典数据并返回分页数据和总行数
func ListUser(params *user_types.UserQuery) ([]*user_types.User, int64, error) {
	var data []*user_types.User
	var count int64
	// 使用事务确保两次查询的数据一致性
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else if tx.Error != nil {
			tx.Rollback()
		}
	}()
	query := tx.Model(&user_types.User{}).Omit("password")
	// 添加字符串字段的模糊匹配条件
	if params.Username != "" {
		query = query.Where("username LIKE ?", "%"+params.Username+"%")
	}
	if params.Email != "" {
		query = query.Where("email LIKE ?", "%"+params.Email+"%")
	}
	if params.Nickname != "" {
		query = query.Where("nickname LIKE ?", "%"+params.Nickname+"%")
	}
	if params.Phone != "" {
		query = query.Where("phone LIKE ?", "%"+params.Phone+"%")
	}

	// 计算总行数
	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	// 获取分页数据
	offset := (params.Page - 1) * params.PageSize
	if err := query.Offset(offset).Limit(params.PageSize).Find(&data).Error; err != nil {
		return nil, 0, err
	}
	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return nil, 0, err
	}
	return data, count, nil
}

// SetAdminUser 设置用户管理员属性
func SetAdminUser(req *user_types.User) error {
	if req.Username == "admin" {
		return fmt.Errorf("禁止修改admin管理员用户")
	}
	result := db.Model(&user_types.User{}).Where("username = ?", req.Username).
		Select("is_admin").
		Updates(req)
	if result.Error != nil {
		return result.Error // 处理错误
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("未找到匹配的记录") // 没有更新任何记录
	}
	return nil // 更新成功
}

// UpdateProfile 更新个人信息
func UpdateProfile(req *user_types.User) error {
	// 列出所有需要更新的字段
	fieldsToUpdate := []string{
		"email", "nickname", "gender", "phone",
		"birthdate", "address", "name", "id_card", "avatar",
	}
	result := db.Model(&user_types.User{}).Where("username = ?", req.Username).
		Select(fieldsToUpdate).
		Updates(req)
	if result.Error != nil {
		return result.Error // 处理错误
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("未找到匹配的记录") // 没有更新任何记录
	}
	return nil // 更新成功
}
