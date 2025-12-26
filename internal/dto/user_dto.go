package dto

import "time"

type RegisterRequest struct {
	Name string `json:"name"`
}

type RegisterResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetAllUsersResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     *string   `json:"email"`
	Age       *int64    `json:"age"`
	CreatedAt time.Time `json:"created_at"`
}

type FindUserByIdResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     *string   `json:"email"`
	Age       *int64    `json:"age"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateUserRequest struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int64  `json:"age"`
}

type UpdateUserResponse struct {
}

type SoftDeleteUserResponse struct {
}

type ForceDeleteUserResponse struct {
}
