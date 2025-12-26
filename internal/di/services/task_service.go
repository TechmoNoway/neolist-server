package di

import (
	"context"
	"neolist-backend/internal/dto"
)

type ITaskSerivce interface {
	Create(ctx context.Context, task dto.CreateTaskRequest) (string, error)
	GetByID(ctx context.Context, is string) (*dto.GetTaskByIDResponse, error)
	Update(ctx context.Context, task dto.UpdateTaskRequest) error
	Patch(ctx context.Context, id string, fields map[string]any) error
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context) ([]*dto.GetAllTasksResponse, error)
	List(ctx context.Context, userID string, opts dto.TaskListOptions) ([]*dto.ListTaskResponse, error)
	Count(ctx context.Context, userID string, opts dto.TaskListOptions) (int64, error)
	Search(ctx context.Context, userID string, query string) ([]*dto.ListTaskResponse, error)
	MarkCompleted(ctx context.Context, id string, completed bool) error
	MarkCompletedMany(ctx context.Context, ids []string, completed bool) error
}
