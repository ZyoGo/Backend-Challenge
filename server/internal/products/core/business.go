package core

import "context"

type Business interface {
	GetProducts(ctx context.Context, dto GetProductsParams) ([]Product, error)
}
