package core

import "time"

type Product struct {
	ID          string
	CategoryID  string
	Name        string
	Description string
	Stock       int
	Price       float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
