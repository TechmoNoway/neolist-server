package services

import (
	"context"
	diRepo "neolist-backend/internal/di/repositories"
	diSvc "neolist-backend/internal/di/services"
	"neolist-backend/internal/dto"
	"neolist-backend/internal/models"

	"github.com/google/uuid"
)

type userService struct {
	userRepo diRepo.IUserRepository
}

func NewUserService(userRepo diRepo.IUserRepository) diSvc.IUserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) Register(ctx context.Context, user dto.RegisterRequest) (*dto.RegisterResponse, error) {

	newUser := &models.UserModel{
		ID:   uuid.New().String(),
		Name: user.Name,
	}

	userRes, err := s.userRepo.Create(ctx, newUser)
	if err != nil {
		return nil, err
	}

	res := &dto.RegisterResponse{
		ID:        userRes.ID,
		Name:      userRes.Name,
		CreatedAt: userRes.CreatedAt,
		UpdatedAt: userRes.UpdatedAt,
	}

	return res, nil
}

func (s *userService) GetAll(ctx context.Context) ([]*dto.GetAllUsersResponse, error) {

	users, err := s.userRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]*dto.GetAllUsersResponse, len(users))

	for i, u := range users {
		responses[i] = &dto.GetAllUsersResponse{
			ID:        u.ID,
			Name:      u.Name,
			Email:     u.Email,
			Age:       u.Age,
			CreatedAt: u.CreatedAt,
		}
	}

	return responses, nil
}

func (s *userService) FindByID(ctx context.Context, id string) (*dto.FindUserByIdResponse, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	response := &dto.FindUserByIdResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Age:       user.Age,
		CreatedAt: user.CreatedAt,
	}

	return response, nil
}

func (s *userService) Update(ctx context.Context, userData dto.UpdateUserRequest) (string, error) {

	modelData := &models.UserModel{
		ID:    userData.ID,
		Name:  userData.Name,
		Email: &userData.Email,
		Age:   &userData.Age,
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
