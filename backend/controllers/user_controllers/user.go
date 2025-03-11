package user_controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"xiaoyun/backend/service/user_service"
	"xiaoyun/backend/types/base_types"
	"xiaoyun/backend/types/user_types"
	"xiaoyun/backend/utils/resp"
	"xiaoyun/backend/validate"
)

// ListUser 列出用户
func ListUser(c *gin.Context) {
	var params user_types.UserQuery
	//确保json请求格式
	if err := c.ShouldBind(&params); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&params, err), nil)
		log.Print(err)
		return
	}
	user, total, err := user_service.ListUser(&params)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Page(c, user, total, &base_types.BasePage{
		Page:     params.Page,
		PageSize: params.PageSize,
	})
}

// SetAdminUser 设置管理员
func SetAdminUser(c *gin.Context) {
	var req user_types.UserAdminReq
	//确保json请求格式
	if err := c.ShouldBind(&req); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&req, err), nil)
		log.Print(err)
		return
	}
	err := user_service.SetAdminUser(&req)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, nil)
}

// CreateUser 创建用户
func CreateUser(c *gin.Context) {
	var req user_types.UserReq
	//确保json请求格式
	if err := c.ShouldBind(&req); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&req, err), nil)
		log.Print(err)
		return
	}
	err := user_service.CreateUser(&req)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, nil)
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	var req base_types.DeleteStringIds
	if err := c.ShouldBind(&req); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&req, err), nil)
		log.Print(err)
		return
	}
	err := user_service.DeleteUser(req.Ids)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, nil)
}

// 更新用户信息
func UpdateUser(c *gin.Context) {
	var req user_types.UserReq
	if err := c.ShouldBind(&req); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&req, err), nil)
		log.Print(err)
		return
	}
	err := user_service.UpdateUser(&req)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, nil)
}
