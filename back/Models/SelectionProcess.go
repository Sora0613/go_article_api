package Models

import "gorm.io/gorm"

// 選考プロセスに関する親

type SelectionProcess struct {
	gorm.Model

	ArticleID  uint        `json:"article_id"`
	Steps      []Step      `gorm:"foreignKey:SelectionProcessID"` // 選考ステップ
	EntrySheet EntrySheet  // エントリーシート
	Interviews []Interview `gorm:"polymorphic:Target"` // 面接情報（1次面接・2次面接）
}

// 選考ステップ

type Step struct {
	gorm.Model
	//エントリーシート → 1次面接 → 2次面接 というふうに記載したい。
	// 下記テーブルの情報から自動生成。
}

type EntrySheet struct {
	gorm.Model
	SelectionProcessID          uint
	QuestionAndSubmissionMethod string `json:"question_and_submission_method"` // 質問（文字数）・提出方法
	Notes                       string `json:"notes"`                          // ES記入時に留意した点
	ResultNotification          string `json:"result_notification"`            // 結果通知方法
}

// 1次面接、2次面接などの情報を入れる。
// 取得するときは配列に入れると個数分回せる？

type Interview struct {
	gorm.Model
	SelectionProcessID uint   `json:"selection_process_id"` // 親テーブルのID
	InterviewType      string `json:"interview_type"`       // 面接の種類（1次面接、2次面接など）
	InterviewerNumber  int    `json:"interviewer_number"`   // 面接官の数
	StudentNumber      int    `json:"student_number"`       // 学生の数
	Venue              string `json:"venue"`                // 会場
	Time               string `json:"time"`                 // 時間
	ProgressMethod     string `json:"progress_method"`      // 進め方
	QuestionContent    string `json:"question_content"`     // 質問内容
	Atmosphere         string `json:"atmosphere"`           // 雰囲気と社員・企業の印象
	InfluentialPoints  string `json:"influential_points"`   // 先行を左右したポイント・その他
	ResultNotification string `json:"result_notification"`  // 結果連絡
}
