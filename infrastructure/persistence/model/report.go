package model

import (
	"reportify-backend/entity"
	"time"
)

type Report struct {
	ID             string
	Body           string
	ReviewBody     *string
	UserID         string
	OrganizationID string
	User           *User
	Organization   *Organization
	Tasks          []*Task
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (m Report) ToEntity() *entity.Report {
	var tasks []entity.Task
	for _, t := range m.Tasks {
		tasks = append(tasks, *t.ToEntity())
	}
	return &entity.Report{
		ID:         m.ID,
		UserID:     m.UserID,
		Body:       m.Body,
		ReviewBody: m.ReviewBody,
		Tasks:      tasks,
		Timestamp:  m.CreatedAt,
	}
}
