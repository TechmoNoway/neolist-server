package mysql

import (
	"context"
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
	return nil, nil
}

func (r *userRepository) FindByID(ctx context.Context, id string) (*models.UserModel, error) {
	return nil, nil
}

func (r *userRepository) Update(ctx context.Context, id string, userData *models.UserModel) (*models.UserModel, error) {
	return nil, nil
}

func (r *userRepository) SoftDelete(ctx context.Context, id string) error {
	return nil
}

func (r *userRepository) ForceDelete(ctx context.Context, id string) error {
	return nil
}
