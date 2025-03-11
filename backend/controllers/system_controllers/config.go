package system_controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"xiaoyun/backend/service/system_service"
	"xiaoyun/backend/types/system_types"
	"xiaoyun/backend/utils/resp"
	"xiaoyun/backend/validate"
)

// GetSysConfig 获取字典数据
func GetSysConfigMap(c *gin.Context) {
	dict, err := system_service.GetSysConfigMap()
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, dict)
}

// UpdateSysConfigMap 保存配置
func UpdateSysConfigMap(c *gin.Context) {
	var data system_types.SysConfigMap
	//确保json请求格式
	if err := c.ShouldBind(&data); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&data, err), nil)
		log.Print(err)
		return
	}
	err := system_service.UpdateSysConfigMap(&data)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, nil)
}

// GetBaseSysConfigMap 获取基本参数
func GetBaseSysConfigMap(c *gin.Context) {
	resp.Ok(c, gin.H{
		"register_enable": system_service.SysConfigMap.SysRegisterEnable,
	})
}
