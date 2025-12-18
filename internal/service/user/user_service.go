package user

import (
	"context"
	"neolist-backend/internal/models"
	"neolist-backend/internal/repositories"

	"github.com/google/uuid"
)

type UserService interface {
	Register(ctx context.Context, user RegisterRequest) (*RegisterResponse, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) Register(ctx context.Context, user RegisterRequest) (*RegisterResponse, error) {

	newUser := &models.UserModel{
		ID:   uuid.New().String(),
		Name: user.Name,
	}

	userRes, err := s.userRepo.Create(ctx, newUser)
	if err != nil {
		return nil, err
	}

	res := &RegisterResponse{
		ID:        userRes.ID,
		Name:      userRes.Name,
		CreatedAt: userRes.CreatedAt,
		UpdatedAt: userRes.UpdatedAt,
	}

	return res, nil
}
