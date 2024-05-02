package request

type CreateOrderRequest struct {
	Products    []ProductRequest `json:"products"`
	CartItemIds []string         `json:"cart_item_ids"`
	UserID      string
}

type ProductRequest struct {
	ID       string `json:"id"`
	Quantity int    `json:"quantity"`
}
