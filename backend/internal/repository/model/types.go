package model

import (
	"time"
)

type Template struct {
	Id         int
	Name       string
	Frequency  string
	DayOfMonth int
	Enabled    bool
}

type Task struct {
	Id          int
	Template    *Template
	Name        string
	DueDate     time.Time
	CompletedAt *time.Time
	Status      string
	CreatedAt   time.Time
}
