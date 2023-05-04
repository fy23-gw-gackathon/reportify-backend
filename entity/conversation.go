package entity

import "time"

// just for test, for store all conversation histories
type Conversation struct {
	ID   string `json:"id"`
	UserID string `json:"userId"`
	IsFromAI bool  `json:"isFromAi"`
	Content string  `json:"content"`
	CreatedAt time.Time  `json:"createdAt"`
}