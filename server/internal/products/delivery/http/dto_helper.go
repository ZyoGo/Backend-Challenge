package http

import (
	"github.com/ZyoGo/Backend-Challange/internal/products/core"
	"github.com/ZyoGo/Backend-Challange/internal/products/delivery/http/request"
)

func NewGetProductsDTO(req *request.GetProductsParams) core.GetProductsParams {
	return core.GetProductsParams{
		CategoryID: req.CategoryID,
	}
}
