package system_types

import (
	"time"
	"xiaoyun/backend/types/base_types"
)

// DictData 数据字典
type DictData struct {
	base_types.BaseType
	DictType    string `json:"dict_type"`
	Value       string `json:"value"`
	LabelName   string `json:"label_name"`
	IsDefault   bool   `json:"is_default"`
	Type        string `gorm:"default:user" json:"type"`
	ElTagType   string `gorm:"default:primary" json:"el_tag_type"`
	ElTagEffect string `gorm:"default:light" json:"el_tag_effect"`
}

// DictType 字典
type DictType struct {
	DictType  string    `gorm:"primaryKey" json:"dict_type"`
	Name      string    `json:"name"`
	Type      string    `gorm:"default:user" json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DictDataReq 数据字典
type DictDataReq struct {
	ID          uint64 `json:"id"`
	DictType    string `json:"dict_type"`
	Value       string `json:"value"`
	LabelName   string `json:"label_name"`
	IsDefault   bool   `json:"is_default"`
	Type        string `json:"type"`
	ElTagType   string `json:"el_tag_type"`
	ElTagEffect string `json:"el_tag_effect"`
}

// DictTypeReq 字典
type DictTypeReq struct {
	DictType string `json:"dict_type" binding:"required"`
	Name     string `json:"name" binding:"required"`
}
