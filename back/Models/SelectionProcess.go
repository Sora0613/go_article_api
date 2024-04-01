package Models

import "gorm.io/gorm"

// 選考プロセス

type SelectionProcess struct {
	gorm.Model

	ArticleID       uint            `json:"article_id"`
	Steps           Steps           `json:"steps"`            // 選考ステップ
	EntrySheet      EntrySheet      `json:"entry_sheet"`      // エントリーシート
	JobFair         JobFair         `json:"job_fair"`         // 説明会
	WrittenTest     WrittenTest     `json:"written_test"`     // 筆記試験・WEBテスト・適性検査
	GroupDiscussion GroupDiscussion `json:"group_discussion"` // グループディスカッション
	OtherSelection  OtherSelection  `json:"other_selection"`  // その他の選考（ジョブ・インターンなど）
	Interviews      []Interview     `json:"interviews"`       // 面接情報（1次面接・2次面接...）
}

type Steps struct {
	gorm.Model

	SelectionProcessID uint   `gorm:"foreignKey:ID" json:"selection_process_id"`
	Steps              string `json:"steps"`
}

type EntrySheet struct {
	gorm.Model

	SelectionProcessID          uint   `gorm:"foreignKey:ID" json:"selection_process_id"`
	QuestionAndSubmissionMethod string `json:"question_and_submission_method"` // 質問（文字数）・提出方法
	Notes                       string `json:"notes"`                          // ES記入時に留意した点
	ResultNotification          string `json:"result_notification"`            // 結果通知方法
}

// 説明会、筆記試験・WEBテスト・適性検査、グループディスカッション、その他の選考（ジョブ・インターンなど) を追加。
// セレクトボックスで選ぶ形となっており、面接のn次面接に関しては、InterviewTypeから取得し、n+1次面接として追加する。

// 説明会
type JobFair struct {
	gorm.Model

	SelectionProcessID uint   `gorm:"foreignKey:ID" json:"selection_process_id"`
	EventDate          string `json:"event_date"` // 開催時期（参加回数）
	Duration           string `json:"duration"`   // 所要時間
	Content            string `json:"content"`    // 内容
	Impression         string `json:"impression"` // 印象・その他
}

// 筆記試験・WEBテスト・適性検査
type WrittenTest struct {
	gorm.Model

	SelectionProcessID uint   `gorm:"foreignKey:ID" json:"selection_process_id"`
	Content            string `json:"content"`             // 試験内容
	Difficulty         string `json:"difficulty"`          // 難易度・対策・ボーダーライン
	ResultNotification string `json:"result_notification"` // 結果連絡
}

// グループディスカッション
type GroupDiscussion struct {
	gorm.Model

	SelectionProcessID uint   `gorm:"foreignKey:ID" json:"selection_process_id"`
	Topic              string `json:"topic"`               // 課されたテーマ・お題
	Duration           string `json:"duration"`            // 所要時間
	Participants       string `json:"participants"`        // 参加者
	Content            string `json:"content"`             // 内容
	ResultNotification string `json:"result_notification"` // 結果連絡
}

// その他の選考（ジョブ・インターンなど）
type OtherSelection struct {
	gorm.Model

	SelectionProcessID uint   `gorm:"foreignKey:ID" json:"selection_process_id"`
	EventData          string `json:"event_date"`          // 開催時期
	DaysOfWork         string `json:"days_of_work"`        // 実施日数・勤務時間
	Content            string `json:"content"`             // 内容
	ParticipantsUniv   string `json:"participants_univ"`   // 参加者の在籍大学
	Treatment          string `json:"treatment"`           // 待遇
	Impression         string `json:"impression"`          // 印象・その他
	ResultNotification string `json:"result_notification"` // 結果連絡
}

// 面接
type Interview struct {
	gorm.Model

	SelectionProcessID uint   `gorm:"foreignKey:ID" json:"selection_process_id"`
	InterviewType      int    `json:"interview_type"`      // 何次面接か
	InterviewerNumber  int    `json:"interviewer_number"`  // 面接官の数
	StudentNumber      int    `json:"student_number"`      // 学生の数
	Venue              string `json:"venue"`               // 会場
	Time               string `json:"time"`                // 時間
	ProgressMethod     string `json:"progress_method"`     // 進め方
	QuestionContent    string `json:"question_content"`    // 質問内容
	Atmosphere         string `json:"atmosphere"`          // 雰囲気と社員・企業の印象
	InfluentialPoints  string `json:"influential_points"`  // 先行を左右したポイント・その他
	ResultNotification string `json:"result_notification"` // 結果連絡
}
