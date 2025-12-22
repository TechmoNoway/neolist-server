package dto

import "time"

type TaskListOptions struct {
	Completed *bool
	DueFrom   *time.Time
	DueTo     *time.Time
	Search    string
	Limit     int64
	Offset    int64
	SortBy    string
}
