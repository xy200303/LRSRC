package article_service

import (
	"xiaoyun/backend/models"
	"xiaoyun/backend/types/article_types"
	"xiaoyun/backend/types/base_types"
)

// ListArticleGroup 列出文章分组
func ListArticleGroup(req *article_types.ArticleGroupQuery) ([]*article_types.ArticleGroup, int64, error) {
	article, count, err := models.ListArticleGroup(req)
	if err != nil {
		return nil, count, err
	}
	return article, count, nil
}

// GetArticleGroup 获取文章分组
func GetArticleGroup(id int64) (*article_types.ArticleGroup, error) {
	group, err := models.GetArticleGroup(id)
	if err != nil {
		return nil, err
	}
	return group, nil
}

// DeleteArticleGroup 删除文章分组
func DeleteArticleGroup(ids base_types.DeleteIntIds) error {
	err := models.DeleteArticleGroup(ids.Ids)
	if err != nil {
		return err
	}
	return nil
}

// UpdateArticleGroup 更新文章分组
func UpdateArticleGroup(req *article_types.ArticleGroup) error {
	err := models.UpdateArticleGroup(req)
	return err
}

// CreateArticleGroup 创建文章分组
func CreateArticleGroup(req *article_types.ArticleGroup) error {
	err := models.CreateArticleGroup(req)
	return err
}

// ListAllArticleGroup 获取所有结果记录
func ListAllArticleGroup() ([]*article_types.ArticleGroup, error) {
	groups, err := models.ListAllArticleGroup()
	if err != nil {
		return nil, err
	}
	return groups, nil
}
