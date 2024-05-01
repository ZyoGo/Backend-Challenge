package postgresql

import (
	"github.com/ZyoGo/Backend-Challange/internal/cart/core"
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

func (row *CartItem) ToCore() core.CartItem {
	return core.CartItem{
		ID:           row.ID,
		CartID:       row.CartID,
		ProductID:    row.ProductID,
		ProductName:  row.ProductName,
		ProductPrice: row.ProductPrice,
		Quantity:     row.Quantity,
		CreatedAt:    row.CreatedAt.Time,
		UpdatedAt:    row.UpdatedAt.Time,
	}
}
