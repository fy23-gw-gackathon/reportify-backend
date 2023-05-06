package model

import "time"

type Report struct {
	ID         string
	Body       string
	ReviewBody string
	UserID     string
	User       *User
	Tasks      []*Task
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
