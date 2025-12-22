package models

import "time"

type TaskModel struct {
	ID                string
	UserID            string
	Title             string
	Description       string
	Completed         bool
	Priority          int64
	StartAt           *time.Time
	DueAt             *time.Time
	RemindedAt        *time.Time
	RecurrenceRule    *string
	RecurrenceEnabled bool
	OrderIndex        int64
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
