package models

import (
	"app/common"
	"time"
)

type User struct {
	ID         int64     `json:"id"`
	FirstName  string    `json:"first_name" validate:"required"`
	Surname    string    `json:"surname" validate:"required"`
	Email      string    `json:"email" validate:"required,email,min=6,max=32"`
	Password   string    `json:"password" validate:"required,min=8,max=32"`
	Permission int8      `json:"permission" validate:"required"`
	IsActive   bool      `json:"is_active"`
	IsAdmin    bool      `json:"is_admin"`
	Phone      string    `json:"phone"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

func (user User) Validate() common.ErrorResponse {
	return common.Validator(&user)
}
