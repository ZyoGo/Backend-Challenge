package business

import (
	"context"

	"github.com/ZyoGo/Backend-Challange/internal/products/core"
)

type ProductBusiness struct {
	repo core.Repository
}

func NewBusiness(repo core.Repository) core.Business {
	return &ProductBusiness{repo}
}

func (s *ProductBusiness) GetProducts(ctx context.Context, dto core.GetProductsParams) ([]core.Product, error) {
	return s.repo.FindProducts(ctx, dto)
}
