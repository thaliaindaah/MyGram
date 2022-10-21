package models

import "time"

type User struct {
	ID        int        `gorm:"primaryKey" json:"id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Email     string     `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required-Your email is required,email-Invalid email format"`
	Username  string     `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required-Your username is required"`
	Password  string     `gorm:"not null" json:"password" form:"password" valid:"required-Your password is required,minstringlength(6)"`
	Age       int        `gorm:"not null;uniqueIndex" json:"age" form:"age" valid:"required-Your age is required,min"`
}
