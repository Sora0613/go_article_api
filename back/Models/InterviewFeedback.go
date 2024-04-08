package Models

import "gorm.io/gorm"

type InterviewFeedback struct {
	gorm.Model

	// 文章を受け取りたいので、typeをTEXTに。
	ArticleID         uint   `gorm:"foreignKey:ArticleID" json:"article_id"`
	MainFocus         string `gorm:"type:TEXT" json:"main_focus"`         // 本選考ではどのようなことが重視されたと思いますか？
	MemorableQuestion string `gorm:"type:TEXT" json:"memorable_question"` // 面接で印象に残った質問をいくつか挙げてください。
	Preparation       string `gorm:"type:TEXT" json:"preparation"`        // 面接時に心がけていたことはありますか？
	KeyPoints         string `gorm:"type:TEXT" json:"key_points"`         // 各選考プロセスで心がけていたことはありますか？
	ResearchSource    string `gorm:"type:TEXT" json:"research_source"`    // どのサイト・書籍を使って企業研究をおこないましたか？
	Impressions       string `gorm:"type:TEXT" json:"impressions"`        // 社員（及び内定者）にどんな印象を持ちましたか？
	AdviceForFuture   string `gorm:"type:TEXT" json:"advice_for_future"`  // これからこの企業を受ける後輩へのアドバイスをお願いします。
}
