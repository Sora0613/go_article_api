package Models

import "gorm.io/gorm"

type Article struct {
	gorm.Model

	Title             Title             `json:"title"`
	Company           Company           `json:"company"`
	SelectionProcess  SelectionProcess  `json:"selection_process"`
	OBVisits          OBVisits          `json:"ob_visits"`
	Offer             Offer             `json:"offer"`
	InterviewFeedback InterviewFeedback `json:"interview_feedback"`
}
