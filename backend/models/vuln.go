package models

import (
	"fmt"
	"time"
	"xiaoyun/backend/types/vuln_types"
)

// CreateVuln 提交漏洞
func CreateVuln(req *vuln_types.Vuln) error {
	//设置为当前时间
	req.CreatedAt = time.Now()
	result := db.Omit("patch", "patch_type").Create(req)
	if result.Error != nil {
		return result.Error // 返回错误
	}
	return nil
}

// ListVuln 查询漏洞
func ListVuln(params *vuln_types.VulnQuery) ([]*vuln_types.VulnResp, int64, error) {
	var data []*vuln_types.VulnResp
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
	query := tx.Model(&vuln_types.Vuln{})
	// 添加字符串字段的模糊匹配条件
	if params.MunaName != "" {
		query = query.Where("muna_name LIKE ?", "%"+params.MunaName+"%")
	}
	if params.Status != "" {
		query = query.Where("status = ?", params.Status)
	}
	if params.MunaDomain != "" {
		query = query.Where("muna_domain LIKE ?", "%"+params.MunaDomain+"%")
	}
	if params.Type != "" {
		query = query.Where("type = ?", params.Type)
	}
	// 计算总行数
	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	// 获取分页数据
	offset := (params.Page - 1) * params.PageSize
	if err := query.Offset(offset).Limit(params.PageSize).Preload("TypeObj").Preload("CateNameObj").Find(&data).Error; err != nil {
		return nil, 0, err
	}
	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return nil, 0, err
	}
	return data, count, nil
}

// UpdateVuln 更新数据
func UpdateVuln(req *vuln_types.Vuln) error {
	// 列出所有需要更新的字段
	fieldsToUpdate := []string{
		"id", "created_by", "created_by",
	}
	result := db.Model(&vuln_types.Vuln{}).Where("id = ?", req.ID).
		Omit(fieldsToUpdate...).
		Updates(req)
	if result.Error != nil {
		return result.Error // 处理错误
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("未找到匹配的记录") // 没有更新任何记录
	}
	return nil // 更新成功
}

// DeleteVuln 删除数据
func DeleteVuln(ids []string) error {
	result := db.Where("id IN (?)", ids).Delete(&vuln_types.Vuln{})
	if result.Error != nil {
		return result.Error // 处理错误
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("未找到匹配的记录") // 没有删除任何记录
	}
	return nil // 删除成功
}

// 获取漏洞信息
func GetVuln(id string) (*vuln_types.VulnResp, error) {
	var data vuln_types.VulnResp
	result := db.Model(&vuln_types.Vuln{}).Where("id = ?", id).Preload("TypeObj").Preload("CateNameObj").First(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return &data, nil
}

// AuditVuln 审核数据
func AuditVuln(req *vuln_types.VulnAuditReq) error {
	// 列出所有需要更新的字段
	fieldsToUpdate := []string{
		"status", "auditor", "audit_opinion", "level", "score",
	}
	result := db.Model(&vuln_types.Vuln{}).Where("id = ?", req.VulnID).
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
