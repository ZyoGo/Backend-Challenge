package core

import "time"

type CartItem struct {
	ID           string
	CartID       string
	ProductID    string
	ProductName  string
	ProductPrice float64
	Amount       float64
	Quantity     int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Cart struct {
	ID        string
	UserID    string
	CartItem  []CartItem
	CreatedAt time.Time
	UpdatedAt time.Time
}
