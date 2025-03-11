package article_controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"xiaoyun/backend/service/article_service"
	"xiaoyun/backend/types/article_types"
	"xiaoyun/backend/types/base_types"
	"xiaoyun/backend/utils/resp"
	"xiaoyun/backend/validate"
)

// ListArticleGroup 列出文章信息
func ListArticleGroup(c *gin.Context) {
	var req article_types.ArticleGroupQuery
	if err := c.ShouldBind(&req); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&req, err), nil)
		log.Print(err)
		return
	}
	group, count, err := article_service.ListArticleGroup(&req)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Page(c, group, count, &base_types.BasePage{
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	return
}

// GetArticleGroup 展示文章分组信息
func GetArticleGroup(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		resp.Resp(c, 400, err.Error(), nil)
		return
	}
	group, err := article_service.GetArticleGroup(id)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, group)
}

// DeleteArticleGroup 删除文章分组信息
func DeleteArticleGroup(c *gin.Context) {
	var ids base_types.DeleteIntIds
	if err := c.ShouldBind(&ids); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(ids, err), nil)
		return
	}
	err := article_service.DeleteArticleGroup(ids)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, nil)
}

// UpdateArticleGroup 更新文章分组
func UpdateArticleGroup(c *gin.Context) {
	var req article_types.ArticleGroup
	if err := c.ShouldBind(&req); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&req, err), nil)
		return
	}
	err := article_service.UpdateArticleGroup(&req)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, nil)
}

// CreateArticleGroup 创建文章分组
func CreateArticleGroup(c *gin.Context) {
	var req article_types.ArticleGroup
	if err := c.ShouldBind(&req); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&req, err), nil)
		return
	}
	err := article_service.CreateArticleGroup(&req)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, nil)
}

// ListAllArticleGroup 获取文章内容
func ListAllArticleGroup(c *gin.Context) {
	articleServices, err := article_service.ListAllArticleGroup()
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, articleServices)
}

// ListAllArticleGroup 前台接口
func ListHomeAllArticleGroup(c *gin.Context) {
	articleServices, err := article_service.ListAllArticleGroup()
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, articleServices)
}
