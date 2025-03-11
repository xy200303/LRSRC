package models

import (
	"fmt"
	"xiaoyun/backend/types/vuln_types"
)

// ListVulnType 查询列表
func ListVulnType(req *vuln_types.VulnTypeQuery) ([]*vuln_types.VulnType, error) {
	var data []*vuln_types.VulnType
	// 使用事务确保两次查询的数据一致性
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	query := tx.Model(&vuln_types.VulnType{})
	if req.ParentID != 0 {
		query.Or("parent_id = ?", req.ParentID)
	}
	if err := query.Find(&data).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return data, nil
}

// CreateVulnType 增加系统漏洞类型
func CreateVulnType(req *vuln_types.VulnType) error {
	result := db.Create(req)
	if result.Error != nil {
		return result.Error // 返回错误
	}
	return nil
}

// UpdateVulnType 更新系统漏洞类型
func UpdateVulnType(req *vuln_types.VulnType) error {
	// 列出所有需要更新的字段
	fieldsToUpdate := []string{
		"cate_type", "type_name", "desc",
	}
	result := db.Model(&vuln_types.VulnType{}).Where("id = ?", req.ID).
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

// DeleteVulnType 删除系统漏洞类型 *目前只能二级删除
func DeleteVulnType(ids []int64) error {
	result := db.Where("id IN (?) or parent_id in (?)", ids, ids).Delete(&vuln_types.VulnType{})
	if result.Error != nil {
		return result.Error // 处理错误
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("未找到匹配的记录") // 没有删除任何记录
	}
	return nil // 删除成功
}
