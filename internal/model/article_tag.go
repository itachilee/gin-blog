package model

import "gorm.io/gorm"

type ArticleTag struct {
	*gorm.Model
	TagID     uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}

func (t ArticleTag) TableName() string {
	return "blog_article_tag"
}
