package models

import (
	"errors"
	"fmt"
	"xiaoyun/backend/types/article_types"
	"xiaoyun/backend/utils"
)

// CreateArticleComment 创建评论
func CreateArticleComment(data *article_types.ArticleComment) error {
	//过滤文本，获取安全的内容
	data.Content = utils.GetSafeContent(data.Content)
	result := db.Create(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// ListArticleComment 列出文章
func ListArticleComment(req *article_types.ArticleCommentQuery) ([]*article_types.ArticleComment, int64, error) {
	var data []*article_types.ArticleComment
	var count int64
	// 使用事务确保两次查询的数据一致性
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	query := tx.Model(&article_types.ArticleComment{})
	// 计算总行数
	if err := query.Count(&count).Error; err != nil {
		tx.Rollback()
		return nil, 0, err
	}
	// 根据查询类型设置排序
	switch req.QueryType {
	case "new":
		query = query.Order("created_at DESC") // 按创建时间倒序，最新的在前
	case "hot":
		query = query.Order("like_count DESC") // 按点赞数倒序，最热门的在前
	default:
		query = query.Order("created_at DESC") // 默认按时间倒序
	}
	query = query.Where("article_id =?", req.ArticleID)
	// 获取分页数据
	offset := (req.Page - 1) * req.PageSize
	if err := query.Offset(offset).Limit(req.PageSize).Find(&data).Error; err != nil {
		tx.Rollback()
		return nil, 0, err
	}
	tx.Commit()
	return data, count, nil
}

// DeleteArticleComment 删除评论
func DeleteArticleComment(ids []int64) error {
	result := db.Where("id IN (?)", ids).Delete(&article_types.ArticleComment{})
	return result.Error
}

// GetArticleComment 查看评论内容
func GetArticleComment(id int64) (*article_types.ArticleComment, error) {
	var comment article_types.ArticleComment
	result := db.Where("id = ?", id).First(&comment)
	return &comment, result.Error
}

func DeleteMyArticleComment(username string, ids []int64) error {
	if len(ids) == 0 {
		return errors.New("评论ID列表不能为空") // 防止误删全部评论
	}
	result := db.Where("created_by = ? AND id IN (?)", username, ids).
		Delete(&article_types.ArticleComment{})
	// 处理记录不存在的情况（按需求决定是否需要报错）
	if result.RowsAffected == 0 {
		return fmt.Errorf("未找到可删除的评论")
	}
	return result.Error
}
