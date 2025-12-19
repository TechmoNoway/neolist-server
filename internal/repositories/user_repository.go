package repositories

import (
	"context"
	"neolist-backend/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.UserModel) (*models.UserModel, error)
	List(ctx context.Context) ([]*models.UserModel, error)
	FindByID(ctx context.Context, id string) (*models.UserModel, error)
	Update(ctx context.Context, userData *models.UserModel) (string, error)
	ForceDelete(ctx context.Context, id string) error
	SoftDelete(ctx context.Context, id string) error
}
