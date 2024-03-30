package Models

import "gorm.io/gorm"

type OBVisits struct {
	gorm.Model
	ArticleID uint   `gorm:"foreignKey:ArticleID" json:"article_id"`
	Visited   string `json:"visited"` // Boolでも良いかも
}
