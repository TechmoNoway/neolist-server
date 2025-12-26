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

type CreateTaskRequest struct {
}

type CreateTaskResponse struct {
}

type UpdateTaskRequest struct {
}

type UpdateTaskResponse struct {
}

type GetTaskByIDResponse struct {
}

type GetAllTasksResponse struct {

}

type ListTaskResponse struct {

}

type SearchTaskRequest struct {

}

type SearchTaskResponse struct {
	
}