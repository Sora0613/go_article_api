package Models

import "gorm.io/gorm"

type OBVisits struct {
	gorm.Model
	Visited string `json:"visited"`
}
