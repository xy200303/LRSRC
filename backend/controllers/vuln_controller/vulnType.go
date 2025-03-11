package vuln_controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"xiaoyun/backend/service/vuln_service"
	"xiaoyun/backend/types/base_types"
	"xiaoyun/backend/types/vuln_types"
	"xiaoyun/backend/utils/resp"
	"xiaoyun/backend/validate"
)

// ListVulnType 查询漏洞类型
func ListVulnType(c *gin.Context) {
	var params vuln_types.VulnTypeQuery
	//确保json请求格式
	if err := c.ShouldBind(&params); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&params, err), nil)
		log.Print(err)
		return
	}
	result, err := vuln_service.ListVulnType(&params)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, result)
}

// DeleteVulnType 删除漏洞类型
func DeleteVulnType(c *gin.Context) {
	var params base_types.DeleteIntIds
	if err := c.ShouldBind(&params); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&params, err), nil)
		log.Print(err)
		return
	}
	err := vuln_service.DeleteVulnType(&params)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, nil)
}

// CreateVulnType 创建漏洞类型
func CreateVulnType(c *gin.Context) {
	var params vuln_types.VulnType
	if err := c.ShouldBind(&params); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&params, err), nil)
		log.Print(err)
		return
	}
	err := vuln_service.CreateVulnType(&params)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, nil)
}

// UpdateVulnType 更新数据漏洞类型
func UpdateVulnType(c *gin.Context) {
	var params vuln_types.VulnType
	if err := c.ShouldBind(&params); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&params, err), nil)
		log.Print(err)
		return
	}
	err := vuln_service.UpdateVulnType(&params)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, nil)
}

// BuildVulnTypeTree 获取漏洞类型级联选择数据
func BuildVulnTypeTree(c *gin.Context) {
	tree, err := vuln_service.BuildVulnTypeTree()
	if err != nil {
		resp.Err(c, err)
	}
	resp.Ok(c, tree)
}
