package system_controllers

import (
	"github.com/gin-gonic/gin"
	"xiaoyun/backend/service/system_service"
	"xiaoyun/backend/utils/resp"
)

// GetSystemStatus 获取系统状态
func GetSystemStatus(c *gin.Context) {
	data := system_service.GetSysStatus()
	resp.Ok(c, data)
}

//获取系统参数
