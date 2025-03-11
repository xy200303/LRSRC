package file_service

import (
	"xiaoyun/backend/models"
	"xiaoyun/backend/types/system_types"
)

// GetFile 读取文件
func GetFile(fileId string) (*system_types.File, error) {
	file, err := models.GetFile(fileId)
	if err != nil {
		return file, err
	}
	return file, nil
}

// SaveFile 创建文件
func SaveFile(file *system_types.File) (*system_types.File, error) {
	createFile, err := models.SaveFile(file)
	if err != nil {
		return nil, err
	}
	return createFile, nil
}

// GetFile 读取文件
func ListFileByIds(fileId []string) ([]*system_types.File, error) {
	file, err := models.ListFileByIds(fileId)
	if err != nil {
		return file, err
	}
	return file, nil
}
