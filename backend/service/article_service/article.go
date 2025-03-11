package article_service

import (
	"fmt"
	"xiaoyun/backend/models"
	"xiaoyun/backend/service/file_service"
	"xiaoyun/backend/types/article_types"
	"xiaoyun/backend/types/base_types"
	"xiaoyun/backend/types/system_types"
)

// GetArticle 获取文章信息
func GetArticle(articleId uint64) (*article_types.ArticleResp, error) {
	article, err := models.GetArticle(articleId)
	if err != nil {
		return nil, fmt.Errorf("获取文章失败: %w", err)
	}
	// 直接返回转换后的结构体
	resp := &article_types.ArticleResp{
		BaseType: article.BaseType,
		Article:  article,
	}
	// 如果有附件，批量获取文件信息
	if len(article.FileList) > 0 {
		resp.FileObjList, err = file_service.ListFileByIds(article.FileList)
		if err != nil {
			return nil, fmt.Errorf("获取文件列表失败: %w", err)
		}
	} else {
		resp.FileList = base_types.StringList{}
		resp.FileObjList = []*system_types.File{}
	}
	return resp, nil
}

// ListArticle 列出文章
func ListArticle(req *article_types.ArticleQuery) ([]*article_types.Article, int64, error) {
	article, i, err := models.ListArticle(req)
	if err != nil {
		return nil, i, err
	}
	return article, i, nil
}

// CreateArticle 创建文章
func CreateArticle(req *article_types.ArticleReq) error {
	err := models.CreateArticle(&article_types.Article{
		CreatedBy:  req.CreatedBy,
		Title:      req.Title,
		Content:    req.Content,
		Desc:       req.Desc,
		TagList:    req.TagList,
		Type:       req.Type,
		FileList:   req.FileList,
		GroupID:    req.GroupID,
		LevelLimit: req.LevelLimit,
		PosterUrl:  req.PosterUrl,
	})
	if err != nil {
		return err
	}
	return nil
}

// 更新文章
func UpdateArticle(req *article_types.ArticleReq) error {
	err := models.UpdateArticle(&article_types.Article{
		BaseType: base_types.BaseType{
			ID: req.ID,
		},
		CreatedBy:  req.CreatedBy,
		Title:      req.Title,
		Content:    req.Content,
		Desc:       req.Desc,
		TagList:    req.TagList,
		Type:       req.Type,
		FileList:   req.FileList,
		GroupID:    req.GroupID,
		LevelLimit: req.LevelLimit,
		PosterUrl:  req.PosterUrl,
	})
	if err != nil {
		return err
	}
	return nil
}

// DeleteArticleGroup 删除文章分组
func DeleteArticle(ids base_types.DeleteIntIds) error {
	err := models.DeleteArticle(ids.Ids)
	if err != nil {
		return err
	}
	return nil
}
