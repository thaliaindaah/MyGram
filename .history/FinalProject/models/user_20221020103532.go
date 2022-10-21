package models

import "time"

type User struct {
	ID        int        `gorm:"primaryKey" json:"id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Email     string
	Username  string
	Password  string
	Age       int
}
