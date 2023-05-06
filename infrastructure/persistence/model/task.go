package model

import "time"

type Task struct {
	ID         string
	Name       string
	ReportID   string
	Report     *Report
	StartedAt  time.Time
	FinishedAt time.Time
}
