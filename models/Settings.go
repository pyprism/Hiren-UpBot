package models

import "time"

type Mailgun struct {
	ID           uint `gorm:"primary_key"`
	UserID       int
	FromEmail    string `gorm:"not null"`
	ToEmail      string `gorm:"not null"`
	Domain       string `gorm:"not null"`
	ApiKey       string `gorm:"not null"`
	PublicApiKey string `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
