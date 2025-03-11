package article_types

import "xiaoyun/backend/types/base_types"

// ArticleGroup 组信息
type ArticleGroup struct {
	base_types.BaseType
	CreatedBy string `gorm:"type:varchar(40)" json:"created_by"` // 创建者
	Name      string `gorm:"type:varchar(40)" json:"name"`       // 名称
	ParentID  int64  `gorm:"parent_id" json:"parent_id"`         // 父组ID
	Desc      string `gorm:"type:varchar(200)" json:"desc"`      // 描述
}

// ArticleGroupQuery 用于查询的ArticleGroup
type ArticleGroupQuery struct {
	base_types.BasePage
	Name      string `json:"name"` // 名称
	ID        int64  `json:"id"`
	CreatedBy string `json:"created_by"`
}
