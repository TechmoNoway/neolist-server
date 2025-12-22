package di

import (
	"context"
	"neolist-backend/internal/dto"
	"neolist-backend/internal/models"
)

type ITaskRepository interface {
	Create(ctx context.Context, task *models.TaskModel) (string, error)
	GetById(ctx context.Context, id string) (*models.TaskModel, error)
	Update(ctx context.Context, task *models.TaskModel) error
	Patch(ctx context.Context, id string, fields map[string]any) error
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context) ([]*models.TaskModel, error)
	List(ctx context.Context, userID string, opts dto.TaskListOptions) ([]*models.TaskModel, error)
	Count(ctx context.Context, userID string, opts dto.TaskListOptions) (int64, error)
	Search(ctx context.Context, userID string, query string) ([]*models.TaskModel, error)
	MarkCompleted(ctx context.Context, id string, completed bool) error
	MarkCompletedMany(ctx context.Context, userID string, ids []string, completed bool) (string, error)
}
