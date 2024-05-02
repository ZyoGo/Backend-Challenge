package business

import (
	"context"
	"fmt"
	"sync"

	"github.com/ZyoGo/Backend-Challange/internal/orders/core"
	"github.com/ZyoGo/Backend-Challange/pkg/derrors"
)

type OrderBusiness struct {
	repo core.Repository
	id   core.Id
	mu   sync.Mutex
}

func NewBusiness(repo core.Repository, id core.Id) core.Business {
	return &OrderBusiness{
		repo: repo,
		id:   id,
		mu:   sync.Mutex{},
	}
}

func (b *OrderBusiness) CreateOrder(ctx context.Context, dto core.CreateOrderDTO) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	orderID := b.id.Generate()
	tx, err := b.repo.BeginTx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	products, err := b.repo.GetProducts(ctx, dto, tx, dto.IsCarts)
	if err != nil {
		return err
	}

	quantityProduct := b.getQuantityProduct(dto.OrderItems)
	if err := b.checkStockOrderItem(products, quantityProduct); err != nil {
		fmt.Println("masuk sini 1 = ", err)
		return err
	}

	amountPrice := b.calculateAmountPriceOfOrders(products, quantityProduct)
	newOrder := b.newOrderHelper(dto, orderID, amountPrice)

	if err := b.repo.CreateOrder(ctx, tx, newOrder); err != nil {
		fmt.Println("masuk sini 2 = ", err)
		return err
	}

	newOrderItem := b.newOrderItemHelper(dto, products, orderID, quantityProduct)
	if err := b.repo.CreateOrderItem(ctx, tx, newOrderItem); err != nil {
		fmt.Println("masuk sini 3 = ", err)
		return err
	}

	if err := b.repo.CreatePaymentVA(ctx, orderID); err != nil {
		fmt.Println("masuk sini 4 = ", err)
		return err
	}

	if dto.IsCarts {
		if err := b.repo.DeleteCartItems(ctx, tx, dto.CartItemID); err != nil {
			fmt.Println("masuk sini 5 = ", err)
			return err
		}
	}

	if err := b.repo.DecreaseStockProduct(ctx, tx, quantityProduct); err != nil {
		fmt.Println("masuk sini 6 = ", err)
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (b *OrderBusiness) getQuantityProduct(orderItems []core.CreateOrderItemDTO) map[string]int {
	var quantityProduct = make(map[string]int)
	for _, item := range orderItems {
		quantityProduct[item.ProductID] = item.Quantity
	}

	return quantityProduct
}

func (b *OrderBusiness) calculateAmountPriceOfOrders(orderItems []core.OrderItem, quantity map[string]int) float64 {
	var amountPrice float64
	for _, item := range orderItems {
		amountPrice += item.ProductPrice * float64(quantity[item.ProductID])
	}
	return amountPrice
}

func (b *OrderBusiness) checkStockOrderItem(orderItems []core.OrderItem, quantity map[string]int) error {
	for _, item := range orderItems {
		if quantity[item.ProductID] > item.ProductStock {
			return derrors.NewErrorf(derrors.ErrorCodeNotFound, "Out of stock for product with id %s", item.ProductID)
		}
	}
	return nil
}

func (b *OrderBusiness) newOrderItemHelper(dto core.CreateOrderDTO, products []core.OrderItem, orderID string, quantityProduct map[string]int) []core.OrderItem {
	orderItems := make([]core.OrderItem, len(dto.OrderItems))

	for i, product := range products {
		orderItems[i] = core.OrderItem{
			ID:           b.id.Generate(),
			OrderID:      orderID,
			ProductID:    product.ProductID,
			ProductPrice: product.ProductPrice,
			Quantity:     quantityProduct[product.ProductID],
		}
	}

	return orderItems
}

func (b *OrderBusiness) newOrderHelper(dto core.CreateOrderDTO, orderID string, amoutPrice float64) core.Order {
	return core.Order{
		ID:            orderID,
		UserID:        dto.UserID,
		PaymentStatus: core.Pending.ToString(),
		Amount:        amoutPrice,
	}
}
