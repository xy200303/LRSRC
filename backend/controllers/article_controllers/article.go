package article_controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"xiaoyun/backend/middleware"
	"xiaoyun/backend/service/article_service"
	"xiaoyun/backend/types/article_types"
	"xiaoyun/backend/types/base_types"
	"xiaoyun/backend/utils/resp"
	"xiaoyun/backend/validate"
)

// ListArticle 列出文章
func ListArticle(c *gin.Context) {
	var req article_types.ArticleQuery
	if err := c.ShouldBind(&req); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&req, err), nil)
		return
	}
	article, i, err := article_service.ListArticle(&req)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Page(c, article, i, &base_types.BasePage{
		Page:     req.Page,
		PageSize: req.PageSize,
	})
}

// CreateArticle 发表文章
func CreateArticle(c *gin.Context) {
	auth, err := middleware.GetCurrentUser(c)
	if err != nil {
		resp.Resp(c, 401, "读取当前用户出错", nil)
		return
	}
	var req article_types.ArticleReq
	req.CreatedBy = auth.User.Username
	if err := c.ShouldBind(&req); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&req, err), nil)
		return
	}
	err = article_service.CreateArticle(&req)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, nil)
}

// 更新文章
func UpdateArticle(c *gin.Context) {
	auth, err := middleware.GetCurrentUser(c)
	if err != nil {
		resp.Resp(c, 401, "读取当前用户出错", nil)
		return
	}
	var req article_types.ArticleReq
	req.CreatedBy = auth.User.Username
	if err := c.ShouldBind(&req); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&req, err), nil)
		return
	}
	err = article_service.UpdateArticle(&req)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, nil)
}

// 删除文章
func DeleteArticle(c *gin.Context) {
	var req base_types.DeleteIntIds
	if err := c.ShouldBind(&req); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&req, err), nil)
		return
	}
	err := article_service.DeleteArticle(req)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, nil)
}

// ListHomeArticle 获取首页文章
func ListHomeArticle(c *gin.Context) {
	var req article_types.ArticleQuery
	if err := c.ShouldBind(&req); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&req, err), nil)
		return
	}
	article, i, err := article_service.ListArticle(&req)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Page(c, article, i, &base_types.BasePage{
		Page:     req.Page,
		PageSize: req.PageSize,
	})
}

// GetHomeArticle 获取文章内容，检查是否有等级查看
func GetArticle(c *gin.Context) {
	auth, err := middleware.GetCurrentUser(c)
	if err != nil {
		resp.Resp(c, 401, "读取当前用户出错", nil)
		return
	}
	articleIdStr := c.Query("id")
	articleId, err := strconv.ParseInt(articleIdStr, 10, 64)
	if err != nil {
		resp.Err(c, err)
		return
	}
	article, err := article_service.GetArticle(uint64(articleId))
	if err != nil {
		resp.Err(c, err)
		return
	}
	//检查等级
	if auth.User.Level < article.LevelLimit && !auth.IsAdmin {
		resp.Err(c, fmt.Errorf("用户等级不够，无法查看当前文章"))
		return
	}
	resp.Ok(c, article)
}
