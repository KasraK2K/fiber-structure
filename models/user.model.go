package models

import (
	"app/common"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID         int64          `gorm:"type:uint;primaryKey;<-:false" json:"id"`
	FirstName  string         `gorm:"type:string;not null;" json:"first_name" validate:"required"`
	Surname    string         `gorm:"type:string;" json:"surname" validate:"required"`
	Email      string         `gorm:"type:string;unique;not null;" json:"email" validate:"required,email,min=6,max=32"`
	Password   string         `gorm:"type:string;check:length(password) >= 8" json:"password" validate:"required,min=8,max=32"`
	Permission string         `gorm:"type:string;" json:"permission" validate:"required"`
	IsActive   bool           `gorm:"type:boolean;" json:"is_active"`
	IsAdmin    bool           `gorm:"type:boolean;" json:"is_admin"`
	Phone      string         `gorm:"type:string;unique;" json:"phone"`
	CreatedAt  time.Time      `gorm:"autoCreateTime:nano;" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime:nano;" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (user User) Validate() common.ErrorResponse {
	return common.Validator(&user)
}
