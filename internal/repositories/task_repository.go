package repositories

import (
	"context"
	"database/sql"
	"log"
	"neolist-backend/internal/db"
	di "neolist-backend/internal/di/repositories"
	"neolist-backend/internal/dto"
	"neolist-backend/internal/models"
)

type taskRepository struct {
	db *db.Database
}

func NewTaskRepository(database *db.Database) di.ITaskRepository {
	return &taskRepository{
		db: database,
	}
}

func (r *taskRepository) Create(ctx context.Context, task *models.TaskModel) (string, error) {
	query := `INSERT INTO Tasks (id, user_id, title, description) VALUES (?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, task.ID, task.UserID, task.Title, task.Description)
	if err != nil {
		return "", err
	}

	return task.ID, nil
}

func (r *taskRepository) GetById(ctx context.Context, id string) (*models.TaskModel, error) {
	query := `
		SELECT id
			, user_id
			, title
			, description
			, completed
			, priority
			, start_at
			, due_at
			, reminded_at
			, recurrence_rule
			, recurrence_enabled
			, order_index
			, created_at
			, updated_at 
		FROM tasks 
		WHERE id = ?
	`
	var t models.TaskModel
	err := r.db.QueryRowContext(ctx, query, id).Scan(&t.ID, &t.UserID, &t.Title, &t.Description, &t.Completed, &t.Priority, &t.StartAt, &t.DueAt, &t.RemindedAt, &t.RecurrenceRule, &t.RecurrenceEnabled, &t.OrderIndex, &t.CreatedAt, &t.UpdatedAt)
	if err == sql.ErrNoRows {
		log.Fatal("Row Not Found")
		return nil, nil
	}

	if err != nil {
		return nil, nil
	}

	return &t, nil
}

func (r *taskRepository) Update(ctx context.Context, task *models.TaskModel) error {
	query := `
		UPDATE tasks
		SET title = ?
			, description = ?
			, completed = ?
			, priority = ?
			, start_at = ?
			, due_at = ?
			, reminded_at = ?
			, recurrence_rule = ?
			, recurrence_enabled = ?
			, order_index = ?
		WHERE id = ?
	`

	result, err := r.db.ExecContext(ctx, query, task.Title, task.Description, task.Completed, task.Priority, task.StartAt, task.DueAt, task.RemindedAt, task.RecurrenceRule, task.RecurrenceEnabled, task.OrderIndex, task.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *taskRepository) Patch(ctx context.Context, id string, fields map[string]any) error {

	return nil
}

func (r *taskRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM tasks where id = ?`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *taskRepository) GetAll(ctx context.Context) ([]*models.TaskModel, error) {

	query := `
		SELECT id
			, user_id
			, title
			, description
			, completed
			, priority
			, start_at
			, due_at
			, reminded_at
			, recurrence_rule
			, recurrence_enabled
			, order_index
			, created_at
			, updated_at 
		FROM tasks 
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tasks := []*models.TaskModel{}

	for rows.Next() {
		var task models.TaskModel
		err := rows.Scan(&task.ID, &task.UserID, &task.Title, &task.Description, &task.Completed, &task.Priority, &task.StartAt, &task.DueAt, &task.RemindedAt, &task.RecurrenceRule, &task.RecurrenceEnabled, &task.OrderIndex, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}

		tasks = append(tasks, &task)

	}

	return tasks, nil
}

func (r *taskRepository) List(ctx context.Context, userID string, opts dto.TaskListOptions) ([]*models.TaskModel, error) {
	return nil, nil
}

func (r *taskRepository) Count(ctx context.Context, userID string, opts dto.TaskListOptions) (int64, error) {
	return 0, nil
}

func (r *taskRepository) Search(ctx context.Context, userID string, query string) ([]*models.TaskModel, error) {
	return nil, nil
}

func (r *taskRepository) MarkCompleted(ctx context.Context, id string, completed bool) error {
	return nil
}

func (r *taskRepository) MarkCompletedMany(ctx context.Context, userID string, ids []string, completed bool) (string, error) {
	return "", nil
}
