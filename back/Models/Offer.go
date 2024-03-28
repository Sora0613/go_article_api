package Models

import "gorm.io/gorm"

type Offer struct {
	gorm.Model

	Offered        string `json:"offered"`
	OfferedAt      string `json:"offered_at"`
	TaskAfterOffer string `json:"task_after_offer"`
	Constraints    string `json:"constraints"`
}