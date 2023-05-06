package entity

type Report struct {
	// 日報レスポンス
	ID string `json:"id"`
	// ユーザID
	UserID string `json:"userId"`
	// 本文
	Body string `json:"body"`
	// レビュー本文
	ReviewBody *string `json:"reviewBody"`
	// 実施したタスクリスト
	Tasks []Task `json:"tasks"`
}
