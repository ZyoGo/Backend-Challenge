package http

import (
	"github.com/ZyoGo/Backend-Challange/internal/orders/core"
	"github.com/ZyoGo/Backend-Challange/internal/orders/delivery/http/request"
)

func NewCreateOrderDTO(req *request.CreateOrderRequest) core.CreateOrderDTO {
	if len(req.CartItemIds) > 0 {
		return core.CreateOrderDTO{
			UserID:     req.UserID,
			CartItemID: req.CartItemIds,
			IsCarts:    true,
		}
	}

	orderItems := make([]core.CreateOrderItemDTO, len(req.Products))
	if len(req.Products) > 0 {
		for i, product := range req.Products {
			orderItems[i] = core.CreateOrderItemDTO{
				ProductID: product.ID,
				Quantity:  product.Quantity,
			}
		}

		return core.CreateOrderDTO{
			UserID:     req.UserID,
			OrderItems: orderItems,
			IsCarts:    false,
		}
	}

	return core.CreateOrderDTO{}
}
