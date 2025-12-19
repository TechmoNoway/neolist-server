package mysql

import (
	"context"
	"database/sql"
	"log"
	"neolist-backend/internal/db"
	"neolist-backend/internal/models"
	"neolist-backend/internal/repositories"
)

type userRepository struct {
	db *db.Database
}

func NewUserRepository(database *db.Database) repositories.UserRepository {
	return &userRepository{
		db: database,
	}
}

func (r *userRepository) Create(ctx context.Context, user *models.UserModel) (*models.UserModel, error) {

	query := `INSERT INTO users (ID, NAME) VALUES (?, ?)`

	_, err := r.db.ExecContext(ctx, query, user.ID, user.Name)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) List(ctx context.Context) ([]*models.UserModel, error) {
	query := `SELECT id, name, email, age, created_at, updated_at FROM users`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []*models.UserModel{}

	for rows.Next() {
		var u models.UserModel
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, &u)
	}

	return users, rows.Err()
}

func (r *userRepository) FindByID(ctx context.Context, id string) (*models.UserModel, error) {
	print(id)
	query := `SELECT id, name, email, age, created_at, updated_at FROM users WHERE id = ?`

	var u models.UserModel
	err := r.db.QueryRowContext(ctx, query, id).Scan(&u.ID, &u.Name, &u.Email, &u.Age, &u.CreatedAt, &u.UpdatedAt)
	if err == sql.ErrNoRows {
		log.Fatal("Row not found")
		return nil, nil
	}

	if err != nil {
		return nil, nil
	}

	return &u, nil
}

func (r *userRepository) Update(ctx context.Context, userData *models.UserModel) (string, error) {
	print(userData.ID)
	query := `UPDATE users SET name = ?, email = ?, age = ? WHERE id = ?`

	result, err := r.db.ExecContext(ctx, query, userData.Name, userData.Email, userData.Age, userData.ID)
	if err != nil {
		return "", err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", err
	}

	if rowsAffected == 0 {
		return "", sql.ErrNoRows
	}

	return userData.ID, nil
}

func (r *userRepository) SoftDelete(ctx context.Context, id string) error {
	query := `
		UPDATE users
		SET deleted_at = NOW()
		WHERE id = ? AND deleted_at IS NULL
	`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return nil
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return nil
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *userRepository) ForceDelete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = ?`

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
