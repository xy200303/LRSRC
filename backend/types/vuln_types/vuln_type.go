package vuln_types

import "xiaoyun/backend/types/base_types"

// VulnType 漏洞类型
type VulnType struct {
	base_types.BaseType
	ParentID uint64 `gorm:"default:0" json:"parent_id"`
	TypeName string `json:"type_name" binding:"required"` // 类型名称
	Desc     string `json:"desc"`                         //类别描述信息
}

// VulnTypeQuery 漏洞类型
type VulnTypeQuery struct {
	base_types.BaseType
	ParentID uint64 `json:"parent_id"` // 类别名称
}
