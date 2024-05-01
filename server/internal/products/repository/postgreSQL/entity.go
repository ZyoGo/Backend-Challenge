package repository

import (
	"github.com/ZyoGo/Backend-Challange/internal/products/core"
	"github.com/jackc/pgx/v5/pgtype"
)

type Product struct {
	ID          string
	CategoryID  string
	Name        string
	Description string
	Stock       int
	Price       float64
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

func FromCore(params core.Product) (Product, error) {
	return Product{
		ID:          params.ID,
		CategoryID:  params.CategoryID,
		Name:        params.Name,
		Description: params.Description,
		Stock:       params.Stock,
		Price:       params.Price,
	}, nil
}

func (row *Product) ToCore() core.Product {
	return core.Product{
		ID:          row.ID,
		CategoryID:  row.CategoryID,
		Name:        row.Name,
		Description: row.Description,
		Stock:       row.Stock,
		Price:       row.Price,
		CreatedAt:   row.CreatedAt.Time,
		UpdatedAt:   row.UpdatedAt.Time,
	}
}
