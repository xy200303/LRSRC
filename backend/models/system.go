package models

import (
	"gorm.io/gorm/clause"
	"xiaoyun/backend/types/system_types"
)

// GetSysConfig 获取系统配置
func GetSysConfig(sysKey string) (*system_types.SysConfig, error) {
	var dict *system_types.SysConfig
	result := db.Where("sys_key = ?", sysKey).First(&dict)
	if result.Error != nil {
		return nil, result.Error
	}
	return dict, nil
}

// UpdateSysConfig 更新系统配置
func UpdateSysConfig(dict *system_types.SysConfig) error {
	result := db.Where("sys_key = ?", dict.SysKey).
		Select("value").
		Updates(dict)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetSysConfigByGroup 获取文件组配置信息
func GetSysConfigByGroup(group string) ([]*system_types.SysConfig, error) {
	var dict []*system_types.SysConfig
	result := db.Where("sys_group = ?", group).Find(&dict)
	if result.Error != nil {
		return nil, result.Error
	}
	return dict, nil
}

// ListSysConfig 获取所有配置项
func ListSysConfig() ([]*system_types.SysConfig, error) {
	var dict []*system_types.SysConfig
	result := db.Find(&dict)
	if result.Error != nil {
		return nil, result.Error
	}
	return dict, nil
}

// UpdateSysConfigBatch 批量更新
func UpdateSysConfigBatch(reqs []*system_types.SysConfig) error {
	result := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "sys_key"}}, // 冲突列
		UpdateAll: false,                              // 不更新所有字段
		DoUpdates: clause.AssignmentColumns([]string{"value"}),
	}).Create(&reqs)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func CreateSysConfigBatch(reqs []*system_types.SysConfig) error {
	result := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "sys_key"}},
		DoNothing: true,
	}).Create(reqs)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
