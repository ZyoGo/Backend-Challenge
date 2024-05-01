package postgresql

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type CartItem struct {
	ID           string
	CartID       string
	ProductID    string
	ProductName  string
	ProductPrice float64
	Amount       float64
	Quantity     int
	CreatedAt    pgtype.Timestamp
	UpdatedAt    pgtype.Timestamp
}

type Cart struct {
	ID        string
	UserID    string
	CartItems []CartItem
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}
