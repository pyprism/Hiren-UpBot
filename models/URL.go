package models

import "time"

type URL struct {
	ID              uint   `gorm:"primary_key"`
	Name            string `gorm:"not null;unique"`
	Url             string `gorm:"not null"`
	UserID          uint
	PollingInterval int64 `gorm:"not null"`
	AlertThreshold  int64 `gorm:"not null"`
	NextRun 		time.Time   // cron next job tracker
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
