package article_types

import (
	"xiaoyun/backend/types/base_types"
)

// ArticleComment 文章评论信息
type ArticleComment struct {
	base_types.BaseType
	ArticleID uint64 `json:"article_id"`                             // 文章ID
	Content   string `json:"content"`                                // 评论内容
	CreatedBy string `json:"created_by"`                             // 用户名称
	Type      string `gorm:"type:varchar(10);default:1" json:"type"` // 评论类型
	Avatar    string `json:"avatar"`
	Nickname  string `json:"nickname"` //用户昵称
	LikeCount uint64 `gorm:"default:0" json:"like_count"`
	Status    string `gorm:"type:varchar(10);default:1" json:"status"`
}

type ArticleCommentReq struct {
	ID        uint64 `json:"id"`
	ArticleID uint64 `json:"article_id"` // 文章ID
	Content   string `json:"content"`    // 评论内容
	Type      string `json:"type"`       // 评论类型
	Status    string `json:"status"`
}

type ArticleCommentResp struct {
	base_types.BaseType
	ArticleID  uint64 `json:"article_id"`                             // 文章ID
	Content    string `json:"content"`                                // 评论内容
	CreatedBy  string `json:"created_by"`                             // 用户名称
	Type       string `gorm:"type:varchar(10).default:1" json:"type"` // 评论类型
	Avatar     string `json:"avatar"`
	LikeCount  uint64 `gorm:"default:0" json:"like_count"`
	Nickname   string `json:"nickname"` //用户昵称
	Status     uint8  `json:"status"`
	ShowDelete bool   `json:"show_delete"`
}

type ArticleCommentQuery struct {
	base_types.BasePage
	ArticleID uint64 `json:"article_id"`
	CreatedBy string `json:"created_by"`
	QueryType string `json:"query_type"`
}
