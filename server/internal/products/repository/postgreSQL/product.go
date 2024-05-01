package repository

import (
	"context"

	"github.com/ZyoGo/Backend-Challange/internal/products/core"
	"github.com/ZyoGo/Backend-Challange/pkg/derrors"
	"github.com/jackc/pgx/v5/pgxpool"
)

const PostgreErrMsg string = "Postgre error"

type postgreSQL struct {
	db *pgxpool.Pool
}

func NewPostgreSQL(db *pgxpool.Pool) *postgreSQL {
	return &postgreSQL{db}
}

func (repo *postgreSQL) FindProducts(ctx context.Context, params core.GetProductsParams) ([]core.Product, error) {
	args := []interface{}{}

	query := `SELECT 
					id, category_id, name, description, 
					stock, price, created_at, updated_at 
				FROM products`

	if params.CategoryID != "" {
		query = query + ` WHERE category_id = $1 `
		args = append(args, &params.CategoryID)
	}

	rows, err := repo.db.Query(ctx, query, args...)
	if err != nil {
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, PostgreErrMsg)
	}
	defer rows.Close()

	var products []core.Product
	for rows.Next() {
		var product Product
		if err := rows.Scan(
			&product.ID, &product.CategoryID, &product.Name,
			&product.Description, &product.Stock, &product.Price,
			&product.CreatedAt, &product.UpdatedAt,
		); err != nil {
			return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, PostgreErrMsg)
		}

		products = append(products, product.ToCore())
	}

	return products, nil
}
