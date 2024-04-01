package Models

import "gorm.io/gorm"

type Offer struct {
	gorm.Model

	ArticleID      uint   `gorm:"foreignKey:ArticleID" json:"article_id"`
	Offered        string `json:"offered"`          // Boolでもいいかも
	OfferedAt      string `json:"offered_at"`       // 内定連絡時期
	TaskAfterOffer string `json:"task_after_offer"` // 内定後の課題
	Constraints    string `json:"constraints"`      // 選考中・内定後の拘束状況
}
