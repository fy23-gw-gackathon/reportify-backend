package entity

import "time"

// Task - 実施したタスク
type Task struct {
	// タスク名
	Name string `json:"name"`
	// 開始日時
	StartedAt time.Time `json:"startedAt"`
	// 終了日時
	FinishedAt time.Time `json:"finishedAt"`
}
