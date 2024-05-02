package core

import "time"

type PaymentStatus int

const (
	Pending PaymentStatus = iota
	Process
	Complete
)

func (s PaymentStatus) ToString() string {
	switch s {
	case Pending:
		return "PENDING"
	case Process:
		return "PROCESS"
	case Complete:
		return "COMPLETE"
	default:
		return "PENDING"
	}
}

type Order struct {
	ID            string
	UserID        string
	PaymentStatus string
	Amount        float64
	CartItemID    []string
	OrderItems    []OrderItem
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type OrderItem struct {
	ID           string
	OrderID      string
	ProductID    string
	ProductName  string
	ProductPrice float64
	ProductStock int
	Quantity     int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
