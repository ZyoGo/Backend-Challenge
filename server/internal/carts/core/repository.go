package core

import "context"

type Repository interface {
	CheckCartByUserID(ctx context.Context, userID string) (exist bool, id string, err error)
	CreateCart(ctx context.Context, userID string, newID string) error
	InsertOrUpdateCartItem(ctx context.Context, params AddCartItemParams) error
	FindCartItems(ctx context.Context, cartID string) (Cart, error)
	DeleteCartItemByID(ctx context.Context, cartItemID string) error
}
