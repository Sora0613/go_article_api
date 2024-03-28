package Models

import "gorm.io/gorm"

type Article struct {
	gorm.Model `json:"gorm_._model"`

	Title      string `json:"title"`
	CompanyID  uint   `json:"company_id"`   // get Company's info from this ID
	OBVisitsID uint   `json:"ob_visits_id"` // get OBVisits's info from this ID
}
