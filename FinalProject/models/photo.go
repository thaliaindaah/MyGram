package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

type Photo struct {
	ID        int        `gorm:"primaryKey" json:"id"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
	Title     string     `gorm:"column:title" json:"title" form:"title" validate:"required"`
	PhotoURL  string     `gorm:"column:photo_url" json:"photo_url" form:"photo_url" validate:"required"`
	Caption   string     `gorm:"column:caption" json:"caption" form:"caption" validate:"required"`
	UserID    int        `gorm:"column:user_id" json:"user_id"`
}

func (Photo) TableName() string {
	return "Photo"
}

func (u *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	vd := validator.New()
	err = vd.Struct(u)
	if err != nil {
		return
	}
	err = nil
	return
}
