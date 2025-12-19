package user

import (
	"context"
	"neolist-backend/internal/models"
	"neolist-backend/internal/repositories"

	"github.com/google/uuid"
)

type UserService interface {
	Register(ctx context.Context, user RegisterRequest) (*RegisterResponse, error)
	List(ctx context.Context) ([]*ListResponse, error)
	FindByID(ctx context.Context, id string) (*FindByIdResponse, error)
	Update(ctx context.Context, userData UpdateRequest) (string, error)
	SoftDelete(ctx context.Context, id string) error
	ForceDelete(ctx context.Context, id string) error
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

func (s *userService) List(ctx context.Context) ([]*ListResponse, error) {

	users, err := s.userRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]*ListResponse, len(users))

	for i, u := range users {
		responses[i] = &ListResponse{
			ID:        u.ID,
			Name:      u.Name,
			Email:     u.Email,
			Age:       u.Age,
			CreatedAt: u.CreatedAt,
		}
	}

	return responses, nil
}

func (s *userService) FindByID(ctx context.Context, id string) (*FindByIdResponse, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	response := &FindByIdResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Age:       user.Age,
		CreatedAt: user.CreatedAt,
	}

	return response, nil
}

func (s *userService) Update(ctx context.Context, userData UpdateRequest) (string, error) {

	modelData := &models.UserModel{
		ID:    userData.ID,
		Name:  userData.Name,
		Email: userData.Email,
		Age:   userData.Age,
	}

	result, err := s.userRepo.Update(ctx, modelData)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (s *userService) SoftDelete(ctx context.Context, id string) error {

	err := s.userRepo.SoftDelete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) ForceDelete(ctx context.Context, id string) error {
	err := s.userRepo.ForceDelete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
