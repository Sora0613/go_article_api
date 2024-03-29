package Models

import "gorm.io/gorm"

type Company struct {
	gorm.Model

	ArticleID         uint   `gorm:"foreignKey:ID" json:"article_id"`
	Name              string `json:"name"`               // 社名
	Department        string `json:"department"`         // 部門、部署
	RecruitmentPeriod string `json:"recruitment_period"` // 選考時期
}
