package core

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Repository interface {
	// Transaction mode
	BeginTx(ctx context.Context) (pgx.Tx, error)

	// Order
	CreateOrder(ctx context.Context, tx pgx.Tx, order Order) error
	CreateOrderItem(ctx context.Context, tx pgx.Tx, orderItems []OrderItem) error

	// Products
	GetProducts(ctx context.Context, params CreateOrderDTO, tx pgx.Tx, isCarts bool) ([]OrderItem, error)
	DecreaseStockProduct(ctx context.Context, tx pgx.Tx, productQuantities map[string]int) error

	// Carts
	DeleteCartItems(ctx context.Context, tx pgx.Tx, cartItemIds []string) error

	// Payment
	CreatePaymentVA(ctx context.Context, orderID string) error
}
