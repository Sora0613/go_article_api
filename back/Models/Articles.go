package Models

import "gorm.io/gorm"

type Article struct {
	gorm.Model

	Title             Title             `json:"title"`
	Company           Company           `json:"company"`
	OBVisits          OBVisits          `json:"ob_visits"`
	Offer             Offer             `json:"offer"`
	InterviewFeedback InterviewFeedback `json:"interview_feedback"`
}

// SelectionProcess  SelectionProcess  `json:"selection_process"`
