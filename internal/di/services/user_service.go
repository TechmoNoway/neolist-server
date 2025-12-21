package di

import (
	"context"
	"neolist-backend/internal/dto"
)

type IUserService interface {
	Register(ctx context.Context, user dto.RegisterRequest) (*dto.RegisterResponse, error)
	GetAll(ctx context.Context) ([]*dto.GetAllResponse, error)
	FindByID(ctx context.Context, id string) (*dto.FindByIdResponse, error)
	Update(ctx context.Context, userData dto.UpdateRequest) (string, error)
	SoftDelete(ctx context.Context, id string) error
	ForceDelete(ctx context.Context, id string) error
}
