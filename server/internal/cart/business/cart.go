package business

import (
	"context"

	"github.com/ZyoGo/Backend-Challange/internal/cart/core"
)

type CartBusiness struct {
	repo core.Repository
	id   core.Id
}

func NewBusiness(repo core.Repository, id core.Id) core.Business {
	return &CartBusiness{repo, id}
}

func (b *CartBusiness) AddCartItem(ctx context.Context, dto core.AddCartItemParams) error {
	cartExist, cartID, err := b.repo.CheckCartByUserID(ctx, dto.UserID)
	if err != nil {
		return err
	}

	if !cartExist {
		cartID = b.id.Generate()
		if err := b.repo.CreateCart(ctx, dto.UserID, cartID); err != nil {
			return err
		}
	}

	insertCartItemParams := core.AddCartItemParams{
		CartItemID: b.id.Generate(),
		CartID:     cartID,
		ProductID:  dto.ProductID,
		Quantity:   dto.Quantity,
	}
	return b.repo.InsertOrUpdateCartItem(ctx, insertCartItemParams)
}

func (b *CartBusiness) GetCartItems(ctx context.Context, userID string) (core.Cart, error) {
	carts, err := b.repo.FindCartItems(ctx, userID)
	if err != nil {
		return carts, err
	}

	for i := range carts.CartItem {
		item := &carts.CartItem[i]
		item.Amount = b.SumAmount(item.Quantity, item.ProductPrice)
	}

	return carts, nil
}

func (b *CartBusiness) SumAmount(quantity int, price float64) float64 {
	return float64(quantity) * price
}
