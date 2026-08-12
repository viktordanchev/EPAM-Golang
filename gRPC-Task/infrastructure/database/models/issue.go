package models

import "time"

type Issue struct {
	ID string `gorm:"primaryKey"`

	Summary     string
	Description string

	Status     string
	Resolution string
	Type       string
	Priority   string

	ProjectID string
	Project   Project

	AssigneeUserID string

	CreatedAt time.Time
	UpdatedAt time.Time
}
