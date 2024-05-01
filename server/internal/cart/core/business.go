package core

import "context"

type Business interface {
	AddCartItem(ctx context.Context, dto AddCartItemParams) error
	GetCartItems(ctx context.Context, cartID string) (Cart, error)
}