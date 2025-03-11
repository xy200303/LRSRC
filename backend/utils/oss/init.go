package oss

import (
	"xiaoyun/backend/types/system_types"
)

func InitOss(sysConfigMap *system_types.SysConfigMap) {
	InitHuaweiOss(sysConfigMap)
	InitAliYunOss(sysConfigMap)
	InitTencentOss(sysConfigMap)
}
