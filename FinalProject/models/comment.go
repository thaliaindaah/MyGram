package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

type Comment struct {
	ID        int        `gorm:"primaryKey" json:"id"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
	Message   string     `gorm:"column:message" json:"message" form:"message" validate:"required"`
	UserID    int        `gorm:"column:user_id" json:"user_id"`
	PhotoID   int        `gorm:"column:photo_id" json:"photo_id"`
	User      []User     `json:"Users"`
	Photo     []Photo    `json:"Photo"`
}

func (Comment) TableName() string {
	return "Comment"
}

func (u *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	vd := validator.New()
	err = vd.Struct(u)
	if err != nil {
		return
	}
	err = nil
	return
}
