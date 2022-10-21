package models

import (
	"FinalProject/helpers"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID        int        `gorm:"primaryKey" json:"id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Email     string     `gorm:"not null,unique_index:email" json:"email" form:"email" validate:"required"`
	Username  string     `gorm:"not null,unique_index:username" json:"username" form:"username" validate:"required"`
	Password  string     `gorm:"not null" json:"password" form:"password" validate:"required,min=6"`
	Age       int        `gorm:"not null;uniqueIndex" json:"age" form:"age" validate:"required,gte=8"`
}

func (User) TableName() string {
	return "Users"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	vd := validator.New()
	err = vd.Struct(u)
	if err != nil {
		return
	}
	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
