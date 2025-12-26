package services

import (
	"context"
	di "neolist-backend/internal/di/repositories"
	diRepo "neolist-backend/internal/di/repositories"
	diSvc "neolist-backend/internal/di/services"
	"neolist-backend/internal/dto"
)

type taskService struct {
	taskRepo di.ITaskRepository
}

func NewTaskService(taskRepo diRepo.ITaskRepository) diSvc.ITaskSerivce {
	return &taskService{}
}

func (s *taskService) Create(ctx context.Context, task dto.CreateTaskRequest) (string, error) {
	return "", nil
}

func (s *taskService) GetByID(ctx context.Context, id string) (*dto.GetTaskByIDResponse, error) {
	return nil, nil
}

func (s *taskService) Update(ctx context.Context, task dto.UpdateTaskRequest) error {
	return nil
}

func (s *taskService) Patch(ctx context.Context, id string, fields map[string]any) error {
	return nil
}

func (s *taskService) Delete(ctx context.Context, id string) error {
	return nil
}

func (s *taskService) GetAll(ctx context.Context) ([]*dto.GetAllTasksResponse, error) {
	return nil, nil
}

func (s *taskService) List(ctx context.Context, userID string, opts dto.TaskListOptions) ([]*dto.ListTaskResponse, error) {
	return nil, nil
}

func (s *taskService) Count(ctx context.Context, userID string, opts dto.TaskListOptions) (int64, error) {
	return 0, nil
}

func (s *taskService) Search(ctx context.Context, userID string, query string) ([]*dto.ListTaskResponse, error) {
	return nil, nil
}

func (s *taskService) MarkCompleted(ctx context.Context, id string, completed bool) error

func (s *taskService) MarkCompletedMany(ctx context.Context, ids []string, completed bool) error
