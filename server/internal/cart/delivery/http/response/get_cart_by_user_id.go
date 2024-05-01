package response

import (
	"net/http"
	"time"

	"github.com/ZyoGo/Backend-Challange/internal/cart/core"
)

type cartItemResp struct {
	ID           string    `json:"id"`
	ProductID    string    `json:"product_id"`
	ProductName  string    `json:"product_name"`
	ProductPrice float64   `json:"product_price"`
	Amount       float64   `json:"amount"`
	Quantity     int       `json:"quantity"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type getCartByUserIDResp struct {
	ID        string         `json:"id"`
	UserID    string         `json:"user_id"`
	CartItems []cartItemResp `json:"cart_items"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type defaultResponse struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Payload getCartByUserIDResp `json:"payload"`
}

func NewGetCartByUserIDResp(cart core.Cart) *defaultResponse {
	var cartItems = make([]cartItemResp, len(cart.CartItem))

	for i, cartItem := range cart.CartItem {
		cartItems[i] = cartItemResp{
			ID:           cartItem.ID,
			ProductID:    cartItem.ProductID,
			ProductName:  cartItem.ProductName,
			ProductPrice: cartItem.ProductPrice,
			Amount:       cartItem.Amount,
			Quantity:     cartItem.Quantity,
			CreatedAt:    cartItem.CreatedAt,
			UpdatedAt:    cartItem.UpdatedAt,
		}
	}

	return &defaultResponse{
		Code:    http.StatusOK,
		Message: "SUCCESS",
		Payload: getCartByUserIDResp{
			ID:        cart.ID,
			UserID:    cart.UserID,
			CartItems: cartItems,
			CreatedAt: cart.CreatedAt,
			UpdatedAt: cart.UpdatedAt,
		},
	}
}
