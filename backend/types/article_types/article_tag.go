package article_types

type ArticleTag struct {
	ArticleID uint64 `gorm:"column:article_id;" json:"article_id"`   //文章ID
	Name      string `json:"name"`                                   //标签名称
	Type      string `gorm:"type:varchar(10);default:1" json:"type"` //标签类型
}
