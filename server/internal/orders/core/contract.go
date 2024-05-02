package core

type CreateOrderDTO struct {
	UserID     string
	CartItemID []string
	OrderItems []CreateOrderItemDTO
	IsCarts    bool
}

type CreateOrderItemDTO struct {
	ProductID string
	Quantity  int
}
