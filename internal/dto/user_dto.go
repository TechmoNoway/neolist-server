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

type GetAllResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     *string   `json:"email"`
	Age       *int64    `json:"age"`
	CreatedAt time.Time `json:"created_at"`
}

type FindByIdResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     *string   `json:"email"`
	Age       *int64    `json:"age"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateRequest struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int64  `json:"age"`
}

type UpdateResponse struct {
}

type SoftDeleteResponse struct {
}

type ForceDeleteResponse struct {
}
