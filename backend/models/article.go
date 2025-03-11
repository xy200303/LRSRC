package models

import (
	"gorm.io/gorm"
	"xiaoyun/backend/types/article_types"
	"xiaoyun/backend/utils"
)

// ListArticle 查询字典类型并返回分页数据和总行数，并支持按标题查询
func ListArticle(page *article_types.ArticleQuery) ([]*article_types.Article, int64, error) {
	var data []*article_types.Article
	var count int64

	// 使用事务确保两次查询的数据一致性
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 构建查询条件
	query := tx.Model(&article_types.Article{})
	if page.Title != "" { // 如果提供了标题，则添加标题过滤条件
		query = query.Where("title LIKE ?", "%"+page.Title+"%")
	}
	if page.GroupID != 0 { // 如果提供了标题，则添加标题过滤条件
		query = query.Where("group_id = ?", page.GroupID)
	}
	if page.Type != "" { // 如果提供了标题，则添加标题过滤条件
		query = query.Where("type = ?", page.Type)
	}

	// 计算总行数
	if err := query.Count(&count).Error; err != nil {
		tx.Rollback()
		return nil, 0, err
	}

	// 获取分页数据
	offset := (page.Page - 1) * page.PageSize
	if err := query.Offset(offset).Limit(page.PageSize).Find(&data).Error; err != nil {
		tx.Rollback()
		return nil, 0, err
	}

	tx.Commit()
	return data, count, nil
}

// GetArticle 获取文章内容
func GetArticle(articleId uint64) (*article_types.Article, error) {
	var data article_types.Article
	// 直接更新浏览量，不使用事务
	if err := db.Model(&article_types.Article{}).Where("id = ?", articleId).UpdateColumn("view_count", gorm.Expr("view_count + 1")).Error; err != nil {
		return nil, err
	}
	// 查询文章内容
	if err := db.Preload("TagList").Where("id = ?", articleId).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

// 添加文章内容
func CreateArticle(data *article_types.Article) error {
	//过滤文本，获取安全的内容
	data.Content = utils.GetSafeContent(data.Content)
	result := db.Create(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 更新文章内容
func UpdateArticle(data *article_types.Article) error {
	// 开始事务
	tx := db.Begin()
	if err := tx.Error; err != nil {
		return err
	}
	// 更新文章内容
	result := tx.Model(&article_types.Article{}).Where("id = ?", data.ID).Updates(&data)
	if result.Error != nil {
		tx.Rollback() // 如果更新失败，回滚事务
		return result.Error
	}
	// 更新TagList（假设有一个关联表 article_tags）
	// 首先删除旧的标签关联
	if err := tx.Where("article_id = ?", data.ID).Delete(&article_types.ArticleTag{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 然后添加新的标签关联
	for _, tag := range data.TagList {
		if err := tx.Create(&article_types.ArticleTag{
			ArticleID: data.ID,
			Name:      tag.Name,
			Type:      tag.Type,
		}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	// 提交事务
	return tx.Commit().Error
}

// 删除文章
func DeleteArticle(articleId []int64) error {
	// 开始事务
	tx := db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	// 删除 ArticleTag 绑定
	if err := tx.Where("article_id IN (?)", articleId).Delete(&article_types.ArticleTag{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 删除文章
	if err := tx.Where("id IN (?)", articleId).Delete(&article_types.Article{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 提交事务
	return tx.Commit().Error
}
