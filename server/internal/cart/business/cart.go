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

func (b *CartBusiness) GetCartItems(ctx context.Context, cartID string) (core.Cart, error) {
	return core.Cart{}, nil
}
