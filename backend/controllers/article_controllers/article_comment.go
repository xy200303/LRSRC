package article_controllers

import (
	"github.com/gin-gonic/gin"
	"xiaoyun/backend/middleware"
	"xiaoyun/backend/service/article_service"
	"xiaoyun/backend/types/article_types"
	"xiaoyun/backend/types/base_types"
	"xiaoyun/backend/utils/resp"
	"xiaoyun/backend/validate"
)

// SendArticleComment 发表评论
func SendArticleComment(c *gin.Context) {
	auth, err := middleware.GetCurrentUser(c)
	if err != nil || auth == nil {
		resp.Resp(c, 401, "读取当前用户出错", nil)
		return
	}
	var req *article_types.ArticleCommentReq
	if err := c.ShouldBind(&req); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&req, err), nil)
		return
	}
	err = article_service.SendArticleComment(&article_types.ArticleComment{
		ArticleID: req.ArticleID,
		Content:   req.Content,
		CreatedBy: auth.User.Username,
		Type:      req.Type,
		Avatar:    auth.User.Avatar,
		Nickname:  auth.User.Nickname,
		Status:    req.Status,
	})
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, req)
	return
}

// ListArticleComment 获取评论数据
func ListArticleComment(c *gin.Context) {
	auth, err := middleware.GetCurrentUser(c)
	if err != nil || auth == nil {
		resp.Resp(c, 401, "读取当前用户出错", nil)
		return
	}
	var req article_types.ArticleCommentQuery
	if err := c.ShouldBind(&req); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&req, err), nil)
		return
	}
	comment, i, err := article_service.ListArticleComment(&req)
	if err != nil {
		resp.Err(c, err)
		return
	}
	//处理评论内容
	var res []*article_types.ArticleCommentResp
	for _, v := range comment {
		commentResp := &article_types.ArticleCommentResp{
			BaseType: base_types.BaseType{
				ID:        v.ID,
				CreatedAt: v.CreatedAt,
			},
			Content:    v.Content,
			CreatedBy:  v.CreatedBy,
			LikeCount:  v.LikeCount,
			ArticleID:  v.ArticleID,
			Avatar:     v.Avatar,
			ShowDelete: (v.CreatedBy == auth.User.Username) || (auth.IsAdmin),
		}
		res = append(res, commentResp)
	}
	resp.Page(c, res, i, &base_types.BasePage{
		Page:     req.Page,
		PageSize: req.PageSize,
	})
}

// DeleteArticleComment 删除评论数据
func DeleteArticleComment(c *gin.Context) {
	var req base_types.DeleteIntIds
	if err := c.ShouldBind(&req); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&req, err), nil)
		return
	}
	err := article_service.DeleteArticleComment(&req)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, nil)
}

// DeleteMyArticleComment 删除自己的评论
func DeleteMyArticleComment(c *gin.Context) {
	auth, err := middleware.GetCurrentUser(c)
	if err != nil || auth == nil {
		resp.Resp(c, 401, "读取当前用户出错", nil)
		return
	}
	var req base_types.DeleteIntIds
	if err := c.ShouldBind(&req); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&req, err), nil)
		return
	}
	err = article_service.DeleteMyArticleComment(auth, &req)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, nil)
}
