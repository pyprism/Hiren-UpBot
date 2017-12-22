package models

import "time"

type User struct {
	ID        uint `gorm:"primary_key"`
	UserName  string    `gorm:"not null;unique"`
	Password  string    `gorm:"not null"`
	Admin     bool      `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
