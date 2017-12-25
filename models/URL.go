package models

import "time"

type URL struct {
	ID              uint   `gorm:"primary_key"`
	Name            string `gorm:"not null;unique"`
	Url             string `gorm:"not null"`
	UserID          int
	PollingInterval string `gorm:"not null"`
	AlertThreshold  string `gorm:"not null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
