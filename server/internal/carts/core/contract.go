package core

type AddCartItemParams struct {
	UserID     string
	CartID     string
	CartItemID string
	ProductID  string
	Quantity   int
}

type DeleteCartItemParams struct {
	UserID     string
	CartItemID string
}
