package models

import (
	"fmt"
	"xiaoyun/backend/types/article_types"
)

// 列出文章分组
func ListArticleGroup(req *article_types.ArticleGroupQuery) ([]*article_types.ArticleGroup, int64, error) {
	var data []*article_types.ArticleGroup
	var count int64
	// 使用事务确保两次查询的数据一致性
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 计算总行数
	if err := tx.Model(&article_types.ArticleGroup{}).Count(&count).Error; err != nil {
		tx.Rollback()
		return nil, 0, err
	}

	// 获取分页数据
	offset := (req.Page - 1) * req.PageSize
	if err := tx.Offset(offset).Limit(req.PageSize).Find(&data).Error; err != nil {
		tx.Rollback()
		return nil, 0, err
	}
	tx.Commit()
	return data, count, nil
}

// GetArticleGroup 获取指定分组信息
func GetArticleGroup(id int64) (*article_types.ArticleGroup, error) {
	var data *article_types.ArticleGroup
	result := db.Where("id = ?", id).First(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}

// DeleteArticleGroup 删除分组信息
func DeleteArticleGroup(ids []int64) error {
	result := db.Where("id IN (?)", ids).Delete(&article_types.ArticleGroup{})
	return result.Error
}

// UpdateArticleGroup 更新数据
func UpdateArticleGroup(req *article_types.ArticleGroup) error {
	// 列出所有需要更新的字段
	fieldsToUpdate := []string{
		"name", "parent_id", "desc",
	}
	result := db.Model(&article_types.ArticleGroup{}).Where("id = ?", req.ID).
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

// CreateArticleGroup 增加新分组信息
func CreateArticleGroup(req *article_types.ArticleGroup) error {
	result := db.Create(req)
	return result.Error
}

// ListAllArticleGroup 获取文章分组DictData
func ListAllArticleGroup() ([]*article_types.ArticleGroup, error) {
	var data []*article_types.ArticleGroup
	err := db.Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, err
}
