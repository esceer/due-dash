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
	TemplateId  *int      `json:"templateId" gorm:"column:template_id"`
	Template    *Template `json:"template" gorm:"foreignKey:TemplateId"`
	Name        string
	DueDate     time.Time
	CompletedAt *time.Time
	Status      string
	CreatedAt   time.Time
}
