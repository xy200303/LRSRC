package models

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"xiaoyun/backend/types/base_types"
	"xiaoyun/backend/types/system_types"
)

// ListDictType 查询字典类型并返回分页数据和总行数
func ListDictType(page *base_types.BasePage) ([]*system_types.DictType, int64, error) {
	var dicts []*system_types.DictType
	var count int64
	// 使用事务确保两次查询的数据一致性
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 计算总行数
	if err := tx.Model(&system_types.DictType{}).Count(&count).Error; err != nil {
		tx.Rollback()
		return nil, 0, err
	}

	// 获取分页数据
	offset := (page.Page - 1) * page.PageSize
	if err := tx.Offset(offset).Limit(page.PageSize).Find(&dicts).Error; err != nil {
		tx.Rollback()
		return nil, 0, err
	}
	tx.Commit()
	return dicts, count, nil
}

// ListDictData 查询字典数据并返回分页数据和总行数
func ListDictData(dictType string, page *base_types.BasePage) ([]*system_types.DictData, int64, error) {
	var dicts []*system_types.DictData
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
	// 定义查询条件
	query := tx.Model(&system_types.DictData{}).Where("dict_type = ?", dictType)
	// 计算总行数
	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	// 获取分页数据
	offset := (page.Page - 1) * page.PageSize
	if err := query.Offset(offset).Limit(page.PageSize).Find(&dicts).Error; err != nil {
		return nil, 0, err
	}
	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return nil, 0, err
	}
	return dicts, count, nil
}

// GetDictData 获取数据字典数据
func GetDictData(dictType string) ([]*system_types.DictData, error) {
	var dataDict []*system_types.DictData
	result := db.Where("dict_type = ?", dictType).Find(&dataDict)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(dataDict) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return dataDict, nil
}

// GetDictType 获取数据字典数据
func GetDictType(dictType string) (*system_types.DictType, error) {
	var dict *system_types.DictType
	result := db.Where("dict_type = ?", dictType).First(&dict)
	if result.Error != nil {
		return nil, result.Error
	}
	return dict, nil
}

// 删除所有DictType所有的数据字典
func DeleteDictDataByDictType(dictTypes []string) error {
	result := db.Where("dict_type IN (?)", dictTypes).Delete(&system_types.DictData{})
	if result.Error != nil {
		return result.Error // 处理错误
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("未找到匹配的记录") // 没有删除任何记录
	}
	return nil // 删除成功
}

// DeleteDictType 删除数据字典数据
func DeleteDictType(dictTypes []string) error {
	result := db.Where("dict_type IN (?)", dictTypes).Delete(&system_types.DictType{})
	if result.Error != nil {
		return result.Error // 处理错误
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("未找到匹配的记录") // 没有删除任何记录
	}
	//删除数据字典
	_ = DeleteDictDataByDictType(dictTypes)
	return nil // 删除成功
}

// DeleteDictData 删除制定键值数据字典
func DeleteDictData(ids []int64) error {
	result := db.Where("id IN (?)", ids).Delete(&system_types.DictData{})
	if result.Error != nil {
		return result.Error // 处理错误
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("未找到匹配的记录") // 没有删除任何记录
	}
	return nil // 删除成功
}

// UpdateDictType 更新数据
func UpdateDictType(req *system_types.DictType) error {
	// 列出所有需要更新的字段
	fieldsToUpdate := []string{
		"name",
	}
	result := db.Model(&system_types.DictType{}).Where("dict_type = ?", req.DictType).
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

// UpdateDictData 更新数据
func UpdateDictData(req *system_types.DictData) error {
	// 列出所有需要更新的字段
	fieldsToUpdate := []string{
		"label_name", "is_default", "el_tag_type", "el_tag_effect",
	}
	result := db.Model(&system_types.DictData{}).Where("id = ?", req.ID).
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

// CreateDictType 创建Dict
func CreateDictType(req *system_types.DictType) error {
	result := db.Create(req)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return fmt.Errorf("当前记录已经存在")
		}
		return result.Error
	}
	return nil
}

// CreateDictData 创建Dict
func CreateDictData(req *system_types.DictData) error {
	result := db.Create(req)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return fmt.Errorf("当前记录已经存在")
		}
		return result.Error
	}
	return nil
}

func ListDictTypeByIds(ids []string) ([]*system_types.DictType, error) {
	var dicts []*system_types.DictType
	result := db.Where("dict_type IN (?)", ids).Find(&dicts)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(dicts) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return dicts, nil
}

func ListDictDataByIds(ids []int64) ([]*system_types.DictData, error) {
	var data []*system_types.DictData
	result := db.Where("id IN (?)", ids).Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(data) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return data, nil
}
