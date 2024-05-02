package repository

import "github.com/ZyoGo/Backend-Challange/internal/orders/core"

type Product struct {
	ID    string
	Name  string
	Price float64
	Stock int
}

func (row *Product) ToCore() core.OrderItem {
	return core.OrderItem{
		ProductID:    row.ID,
		ProductName:  row.Name,
		ProductPrice: row.Price,
		ProductStock: row.Stock,
	}
}
