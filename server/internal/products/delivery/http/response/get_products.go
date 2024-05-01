package response

import (
	"net/http"
	"time"

	"github.com/ZyoGo/Backend-Challange/internal/products/core"
)

type Product struct {
	ID          string    `json:"id"`
	CategoryID  string    `json:"category_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Stock       int       `json:"stock"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetProductsResp struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Payload []Product `json:"payload"`
}

func NewGetProductsResp(products []core.Product) *GetProductsResp {
	productsResponse := make([]Product, len(products))

	for i, product := range products {
		productsResponse[i] = Product{
			ID:          product.ID,
			CategoryID:  product.CategoryID,
			Name:        product.Name,
			Description: product.Description,
			Stock:       product.Stock,
			Price:       product.Price,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
		}
	}

	return &GetProductsResp{
		Code:    http.StatusOK,
		Message: "SUCCESS",
		Payload: productsResponse,
	}
}
