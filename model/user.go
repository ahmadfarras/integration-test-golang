package model

import "time"

type User struct {
	Id        string
	FullName  string
	Password  string
	Email     string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
