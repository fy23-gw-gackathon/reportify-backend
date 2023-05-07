package entity

import "time"

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
	// 作成日時
	Timestamp time.Time `json:"timestamp"`
}

type PubSubMessage struct {
	ID   string `json:"id"`
	Body string `json:"body"`
	Mvv  `json:"mvv"`
}

type ReviewReportRequest struct {
	ReviewBody string `json:"reviewBody"`
}
