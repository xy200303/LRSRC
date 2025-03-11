package vuln_controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"xiaoyun/backend/middleware"
	"xiaoyun/backend/service/vuln_service"
	"xiaoyun/backend/types/base_types"
	"xiaoyun/backend/types/vuln_types"
	"xiaoyun/backend/utils/resp"
	"xiaoyun/backend/validate"
)

// ListVuln 查询漏洞数据
func ListVuln(c *gin.Context) {
	var params vuln_types.VulnQuery
	//确保json请求格式
	if err := c.ShouldBind(&params); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&params, err), nil)
		log.Print(err)
		return
	}
	result, total, err := vuln_service.ListVuln(&params)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Page(c, result, total, &base_types.BasePage{Page: params.Page, PageSize: params.PageSize})
}

// CreateVuln 创建漏洞
func CreateVuln(c *gin.Context) {
	auth, err := middleware.GetCurrentUser(c)
	if err != nil {
		resp.Resp(c, 401, "读取当前用户出错", nil)
		return
	}
	var params vuln_types.Vuln
	if err = c.ShouldBind(&params); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&params, err), nil)
		log.Print(err)
		return
	}
	params.CreatedBy = auth.User.Username
	err = vuln_service.CreateVuln(&params)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, nil)
}

// UpdateVuln 更新漏洞
func UpdateVuln(c *gin.Context) {
	var params vuln_types.Vuln
	if err := c.ShouldBind(&params); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&params, err), nil)
		log.Print(err)
		return
	}
	err := vuln_service.UpdateVuln(&params)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, nil)
}

// DeleteVuln 删除漏洞
func DeleteVuln(c *gin.Context) {
	var ids base_types.DeleteStringIds
	if err := c.ShouldBind(&ids); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&ids, err), nil)
		log.Print(err)
		return
	}
	err := vuln_service.DeleteVuln(&ids)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, nil)
}

// GetVuln 获取漏洞信息
func GetVuln(c *gin.Context) {
	id := c.Query("id")
	vuln, err := vuln_service.GetVuln(id)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, vuln)
}

// SubmitVuln 提交漏洞-用户侧接口
func SubmitVuln(c *gin.Context) {
	auth, err := middleware.GetCurrentUser(c)
	if err != nil {
		resp.Resp(c, 401, "读取当前用户出错", nil)
		return
	}
	var params vuln_types.Vuln
	if err = c.ShouldBind(&params); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&params, err), nil)
		log.Print(err)
		return
	}
	params.CreatedBy = auth.User.Username
	err = vuln_service.SubmitVuln(&params)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, nil)
}

// AuditVuln 审核漏洞-管理员侧接口
func AuditVuln(c *gin.Context) {
	auth, err := middleware.GetCurrentUser(c)
	if err != nil {
		resp.Resp(c, 401, "读取当前用户出错", nil)
		return
	}
	var params vuln_types.VulnAuditReq
	if err = c.ShouldBind(&params); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&params, err), nil)
		log.Print(err)
		return
	}
	params.Auditor = auth.User.Username
	err = vuln_service.AuditVuln(&params)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, nil)
}
