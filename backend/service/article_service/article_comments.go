package article_service

import (
	"fmt"
	"xiaoyun/backend/models"
	"xiaoyun/backend/types"
	"xiaoyun/backend/types/article_types"
	"xiaoyun/backend/types/base_types"
)

func SendArticleComment(req *article_types.ArticleComment) error {
	//判断文章是否存在
	_, err := models.GetArticle(req.ArticleID)
	if err != nil {
		return fmt.Errorf("不存在的文章信息")
	}
	err = models.CreateArticleComment(req)
	if err != nil {
		return err
	}
	return nil
}

// ListArticleComment 列出所有文章
func ListArticleComment(req *article_types.ArticleCommentQuery) ([]*article_types.ArticleComment, int64, error) {
	comment, i, err := models.ListArticleComment(req)
	if err != nil {
		return nil, 0, nil
	}
	return comment, i, nil
}

// DeleteArticleComment 删除文章评论
func DeleteArticleComment(req *base_types.DeleteIntIds) error {
	err := models.DeleteArticleComment(req.Ids)
	return err
}

// DeleteMyArticleComment 删除自己的文章评论
func DeleteMyArticleComment(auth *types.Auth, req *base_types.DeleteIntIds) error {
	var err error
	if auth.IsAdmin {
		err = models.DeleteArticleComment(req.Ids)
	} else {
		err = models.DeleteMyArticleComment(auth.User.Username, req.Ids)
	}
	return err
}
