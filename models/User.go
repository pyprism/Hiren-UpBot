package models

import "time"

type User struct {
	Id        int64
	UserName  string    `xorm:"unique varchar(25) not null"`
	Password  string    `xorm:"varchar(255) not null"`
	Admin     bool      `xorm:"not null"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}
