package http

import (
	"github.com/ZyoGo/Backend-Challange/internal/cart/core"
	"github.com/ZyoGo/Backend-Challange/internal/cart/delivery/http/request"
)

func NewAddCartItemDTO(req *request.AddCartItemRequest) core.AddCartItemParams {
	return core.AddCartItemParams{
		UserID:    req.UserID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}
}
