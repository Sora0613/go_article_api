package Models

import "gorm.io/gorm"

type Title struct {
	gorm.Model
	ArticleID uint   `gorm:"foreignKey:ID" json:"article_id"`
	Title     string `json:"title"`
}
