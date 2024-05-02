package request

type AddCartItemRequest struct {
	UserID    string
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
