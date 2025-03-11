package article_types

import (
	"xiaoyun/backend/types/base_types"
	"xiaoyun/backend/types/system_types"
)

// Article 文章信息
type Article struct {
	base_types.BaseType
	CreatedBy   string                `json:"created_by"` // 用户名称
	Title       string                `gorm:"type:varchar(100)" json:"title"`
	Content     string                `json:"content"`                              // 内容
	Desc        string                `gorm:"type:varchar(256)" json:"desc"`        // 描述
	TagList     []ArticleTag          `gorm:"foreignKey:ArticleID" json:"tag_list"` // 标签
	LikeCount   int                   `gorm:"default:0" json:"like_count"`          // 点赞数
	ViewCount   int                   `gorm:"default:0" json:"view_count"`          // 查看数
	UnlikeCount int                   `gorm:"default:0" json:"unlike_count"`        // 阅读数
	Type        string                `gorm:"type:varchar(10)" json:"type"`         // 类型
	FileList    base_types.StringList `gorm:"type:json" json:"file_list"`           // 文件列表
	GroupID     int                   `json:"group_id"`                             // 组ID
	LevelLimit  uint8                 `json:"level_limit"`                          // 等级限制
	PosterUrl   string                `json:"poster_url"`
}

// ArticleReq 创建文章的记录
type ArticleReq struct {
	base_types.BaseType
	CreatedBy  string                `json:"created_by"` // 用户名称
	Title      string                `json:"title"`
	Content    string                `json:"content"`     // 内容
	Desc       string                `json:"desc"`        // 描述
	TagList    []ArticleTag          `json:"tag_list"`    // 标签
	Type       string                `json:"type"`        // 类型
	FileList   base_types.StringList `json:"file_list"`   // 文件列表
	GroupID    int                   `json:"group_id"`    // 组ID
	LevelLimit uint8                 `json:"level_limit"` // 等级限制
	PosterUrl  string                `json:"poster_url"`
}

// ArticleResp 创建文章的记录
type ArticleResp struct {
	base_types.BaseType
	*Article
	FileObjList []*system_types.File `json:"file_obj_list"` //文件对象
}

// ArticleQuery 文章搜素
type ArticleQuery struct {
	base_types.BasePage
	Title   string                `json:"title"`
	TagList base_types.StringList `json:"tag_list"`
	GroupID int                   `json:"group_id"`
	Type    string                `json:"type"`
}
