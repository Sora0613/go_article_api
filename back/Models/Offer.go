package Models

import "gorm.io/gorm"

type Offer struct {
	gorm.Model

	ArticleID      uint   `gorm:"foreignKey:ArticleID" json:"article_id"`
	Offered        string `json:"offered"` // Boolでもいいかも
	OfferedAt      string `json:"offered_at"`
	TaskAfterOffer string `json:"task_after_offer"`
	Constraints    string `json:"constraints"`
}
