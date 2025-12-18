package models

import "time"

type UserModel struct {
	ID        string
	Name      string
	Email     string
	Age       int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
