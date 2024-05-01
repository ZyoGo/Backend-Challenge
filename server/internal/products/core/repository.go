package core

import "context"

type Repository interface {
	FindProducts(ctx context.Context, params GetProductsParams) ([]Product, error)
}
