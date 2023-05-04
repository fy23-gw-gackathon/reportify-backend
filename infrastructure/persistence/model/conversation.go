package model

import "time"

type Conversation struct {
	ID   string `gorm:"primaryKey"`
	UserID string
	IsFromAI bool 
	Content string
	CreatedAt time.Time  `gorm:"autoCreateTime"`
}