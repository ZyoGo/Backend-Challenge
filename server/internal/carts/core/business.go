package core

import "context"

type Business interface {
	AddCartItem(ctx context.Context, dto AddCartItemParams) error
	GetCartItems(ctx context.Context, userID string) (Cart, error)
	DeleteCartItemByID(ctx context.Context, dto DeleteCartItemParams) error
}
