package core

import "time"

type User struct {
	ID          string
	Username    string
	Email       string
	Password    string
	Address     string
	PhoneNumber string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
