package model

import (
	"time"
)

type Status string

const (
	StatusPending   Status = "PENDING"
	StatusCompleted Status = "COMPLETED"
	StatusOverdue   Status = "OVERDUE"
	StatusSkipped   Status = "SKIPPED"
)

type Template struct {
	Id      int  `json:"id"`
	Enabled bool `json:"enabled"`
	NewTemplate
}

type NewTemplate struct {
	Name       string `json:"name"`
	Frequency  string `json:"frequency"`
	DayOfMonth int    `json:"dayOfMonth"`
}

type Task struct {
	Id          int        `json:"id"`
	Template    *Template  `json:"template"`
	CompletedAt *time.Time `json:"completedAt"`
	CreatedAt   time.Time  `json:"createdAt"`
	Status      Status     `json:"status"`
	NewTask
}

type NewTask struct {
	Name    string    `json:"name"`
	DueDate time.Time `json:"dueDate"`
}
