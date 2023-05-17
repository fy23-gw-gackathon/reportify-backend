package model

import (
	"github.com/fy23-gw-gackathon/reportify-backend/entity"
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
	var userName string
	if m.User != nil {
		userName = m.User.Name
	}
	return &entity.Report{
		ID:         m.ID,
		UserID:     m.UserID,
		UserName:   userName,
		Body:       m.Body,
		ReviewBody: m.ReviewBody,
		Tasks:      tasks,
		Timestamp:  m.CreatedAt,
	}
}
