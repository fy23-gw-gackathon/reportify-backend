package model

import (
	"reportify-backend/entity"
	"time"
)

type Task struct {
	ID         string
	Name       string
	ReportID   string
	Report     *Report
	StartedAt  time.Time
	FinishedAt time.Time
}

func (m Task) ToEntity() *entity.Task {
	return &entity.Task{
		Name:       m.Name,
		StartedAt:  m.StartedAt,
		FinishedAt: m.FinishedAt,
	}
}
