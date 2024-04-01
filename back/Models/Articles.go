package Models

import "gorm.io/gorm"

type Article struct {
	gorm.Model

	Title             Title             `json:"title"`              // 1つの記事に1つのタイトル
	Company           Company           `json:"company"`            // エントリーした企業について
	SelectionProcess  SelectionProcess  `json:"selection_process"`  // 選考について
	OBVisits          OBVisits          `json:"ob_visits"`          // OB訪問について
	Offer             Offer             `json:"offer"`              // 内定後について
	InterviewFeedback InterviewFeedback `json:"interview_feedback"` //面接についての感想
}
