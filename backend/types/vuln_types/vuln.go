package vuln_types

import (
	"gorm.io/gorm"
	"time"
	"xiaoyun/backend/types/base_types"
	"xiaoyun/backend/utils"
)

// Vuln 漏洞数据
type Vuln struct {
	ID               string                `gorm:"primaryKey;" json:"id"` // 漏洞ID
	MunaName         string                `json:"muna_name"`             // 厂商名称
	Desc             string                `json:"desc"`                  // 漏洞描述
	CreatedBy        string                `json:"created_by"`            // 提交漏洞的用户名
	CreatedAt        time.Time             `json:"created_at"`
	UpdatedAt        time.Time             `json:"updated_at"`
	Patch            string                `json:"patch"`                                               // 补丁信息
	PatchType        string                `json:"patch_type"`                                          // 补丁类型
	Title            string                `json:"title"`                                               // 漏洞标题
	MunaDomain       string                `json:"muna_domain"`                                         // 厂商主域
	Detail           string                `json:"detail"`                                              // 漏洞详情
	Type             uint64                `json:"type"`                                                // 漏洞类型
	Level            string                `json:"level"`                                               // 漏洞等级
	AttachmentID     string                `json:"attachment_id"`                                       // 附件ID
	AttachmentName   string                `json:"attachment_name"`                                     // 附件名称
	Status           string                `gorm:"default:under_review;type:varchar(20)" json:"status"` // 漏洞状态
	RepairSuggestion string                `json:"repair_suggestion"`                                   // 修复建议
	Province         string                `json:"province"`                                            // 省份
	City             string                `json:"city"`                                                // 城市
	County           string                `json:"county"`                                              // 区县
	Industry         base_types.StringList `gorm:"type:json" json:"industry"`                           //行业信息
	URL              string                `json:"url"`                                                 // 漏洞URL
	Poc              string                `json:"poc"`                                                 // 漏洞POC
	Score            float32               `json:"score"`                                               //漏洞评分
	Attribute        string                `json:"attribute"`                                           //漏洞属性 通用性或者事件型
	CateType         uint64                `json:"cate_type"`                                           //漏洞类别
	AuditOpinion     string                `json:"audit_opinion"`                                       //审核意见
	Auditor          string                `json:"auditor"`                                             //审核员
}

// VulnResp 返回给前端数据
type VulnResp struct {
	Vuln
	TypeObj     VulnType `gorm:"foreignKey:Type" json:"type_obj"`
	CateNameObj VulnType `gorm:"foreignKey:CateType" json:"cate_name_obj"`
}

// BeforeCreate 钩子函数，在创建记录前自动为 ID 字段生成 UUID
func (v *Vuln) BeforeCreate(tx *gorm.DB) (err error) {
	v.ID = utils.GenerateLRID("LRVE")
	return
}

// VulnAuditReq 审核漏洞
type VulnAuditReq struct {
	VulnID       string  `json:"vuln_id"`
	AuditOpinion string  `json:"audit_opinion"`
	Auditor      string  `json:"auditor"`
	Level        string  `json:"level"`                                               // 漏洞等级
	Status       string  `gorm:"default:under_review;type:varchar(20)" json:"status"` // 漏洞状态
	Score        float32 `json:"score"`                                               //漏洞评分
}

// VulnQuery 查询数据
type VulnQuery struct {
	base_types.BasePage
	MunaName   string `json:"muna_name"`   // 厂商名称
	Level      string `json:"level"`       // 漏洞等级
	Type       string `json:"type"`        // 漏洞类型
	Status     string `json:"status"`      // 漏洞状态
	MunaDomain string `json:"muna_domain"` // 厂商主域
}
