package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

type SocialMedia struct {
	ID     int    `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"column:name" json:"name" form:"name" validate:"required"`
	URL    string `gorm:"column:social_media_url" json:"social_media_url" form:"social_media_url" validate:"required"`
	UserID int    `gorm:"column:user_id" json:"user_id"`
	User   []User `json:"Users"`
}

func (SocialMedia) TableName() string {
	return "SocialMedia"
}

func (u *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	vd := validator.New()
	err = vd.Struct(u)
	if err != nil {
		return
	}
	err = nil
	return
}
