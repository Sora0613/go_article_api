package Models

import "gorm.io/gorm"

// TODO:これらの項目は全て改行を許可したい。

type InterviewFeedback struct {
	gorm.Model

	ArticleID         uint   `gorm:"foreignKey:ArticleID" json:"article_id"`
	MainFocus         string `json:"main_focus"`         // 本選考ではどのようなことが重視されたと思いますか？
	MemorableQuestion string `json:"memorable_question"` // 面接で印象に残った質問をいくつか挙げてください。
	Preparation       string `json:"preparation"`        // 面接時に心がけていたことはありますか？
	KeyPoints         string `json:"key_points"`         // 各選考プロセスで心がけていたことはありますか？
	ResearchSource    string `json:"research_source"`    // どのサイト・書籍を使って企業研究をおこないましたか？
	Impressions       string `json:"impressions"`        // 社員（及び内定者）にどんな印象を持ちましたか？
	AdviceForFuture   string `json:"advice_for_future"`  // これからこの企業を受ける後輩へのアドバイスをお願いします。
}
