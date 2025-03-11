package models

import (
	"fmt"
	"xiaoyun/backend/types/system_types"
)

// GetFile 获取文件对象
func GetFile(fileId string) (*system_types.File, error) {
	var data *system_types.File
	result := db.Where("file_id = ?", fileId).First(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}

// SaveFile 增加记录
func SaveFile(file *system_types.File) (*system_types.File, error) {
	result := db.Save(file) // Save 会自动判断是插入还是更新
	if result.Error != nil {
		return nil, result.Error
	}
	return file, nil
}

// 列出文件
func ListFileByIds(fileIds []string) ([]*system_types.File, error) {
	var data []*system_types.File
	result := db.Where("file_id in (?)", fileIds).Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("file ids is empty")
	}
	return data, nil
}
