package models

import (
	"time"
)

type User struct {
	ID        int
	Name      string `sql:",unique,notnull,type:varchar(255)"`
	Password  string `sql:",notnull,type:varchar(500)"`
	Admin     bool   `sql:",notnull"`
	CreatedAt time.Time
}
