package http

import (
	"github.com/ZyoGo/Backend-Challange/internal/carts/core"
	"github.com/ZyoGo/Backend-Challange/internal/carts/delivery/http/request"
)

func NewAddCartItemDTO(req *request.AddCartItemRequest) core.AddCartItemParams {
	return core.AddCartItemParams{
		UserID:    req.UserID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}
}

func NewDeleteCartItemDTO(req *request.DeleteCartItemReq) core.DeleteCartItemParams {
	return core.DeleteCartItemParams{
		UserID:     req.UserID,
		CartItemID: req.CartItemID,
	}
}
