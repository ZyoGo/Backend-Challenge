package repository

import (
	"context"

	"github.com/ZyoGo/Backend-Challange/internal/orders/core"
	"github.com/ZyoGo/Backend-Challange/pkg/derrors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const PostgreErrMsg string = "Postgre Error"

type PostgreSQL struct {
	db *pgxpool.Pool
}

func NewPostgreSQL(db *pgxpool.Pool) *PostgreSQL {
	return &PostgreSQL{db}
}

func (repo *PostgreSQL) BeginTx(ctx context.Context) (pgx.Tx, error) {
	tx, err := repo.db.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		return tx, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, PostgreErrMsg)
	}

	return tx, nil
}

func (repo *PostgreSQL) getProductsByCartItems(ctx context.Context, tx pgx.Tx, productIds []string) ([]string, error) {
	query := `SELECT product_id FROM cart_items WHERE id = ANY($1)`

	rows, err := tx.Query(ctx, query, productIds)
	if err != nil {
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, PostgreErrMsg)
	}
	defer rows.Close()

	var productsId []string
	for rows.Next() {
		var id string

		err := rows.Scan(&id)
		if err != nil {
			return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, PostgreErrMsg)
		}

		productsId = append(productsId, id)
	}

	return productsId, nil
}

func (repo *PostgreSQL) GetProducts(ctx context.Context, params core.CreateOrderDTO, tx pgx.Tx, isCarts bool) ([]core.OrderItem, error) {
	var (
		productIds []string
		err        error
	)

	if isCarts {
		productIds, err = repo.getProductsByCartItems(ctx, tx, productIds)
		if err != nil {
			return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, PostgreErrMsg)
		}
	} else {
		for _, dto := range params.OrderItems {
			productIds = append(productIds, dto.ProductID)
		}
	}

	query := `SELECT id, name, price, stock FROM products WHERE id = ANY($1) FOR UPDATE`
	rows, err := tx.Query(ctx, query, productIds)
	if err != nil {
		return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, PostgreErrMsg)
	}
	defer rows.Close()

	var products []core.OrderItem
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Stock); err != nil {
			return nil, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, PostgreErrMsg)
		}
		products = append(products, product.ToCore())
	}

	return products, nil
}

func (repo *PostgreSQL) CreateOrder(ctx context.Context, tx pgx.Tx, order core.Order) error {
	query := `INSERT INTO orders(id, user_id, payment_status, amount, created_at, updated_at)
				VALUES($1, $2, $3, $4, now(), now())`

	_, err := tx.Exec(ctx, query, order.ID, order.UserID, order.PaymentStatus, order.Amount)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, PostgreErrMsg)
	}
	return nil
}

func (repo *PostgreSQL) CreateOrderItem(ctx context.Context, tx pgx.Tx, orderItems []core.OrderItem) error {
	batch := &pgx.Batch{}

	for _, item := range orderItems {
		query := `INSERT INTO order_items(id, order_id, product_id, quantity, price, created_at, updated_at)
					VALUES($1, $2, $3, $4, $5, now(), now())`
		batch.Queue(query, item.ID, item.OrderID, item.ProductID, item.Quantity, item.ProductPrice)
	}

	result := tx.SendBatch(ctx, batch)
	defer result.Close()

	for i := 0; i < batch.Len(); i++ {
		result, _ := result.Query()
		if result.Err() != nil {
			return derrors.WrapErrorf(result.Err(), derrors.ErrorCodeUnknown, PostgreErrMsg)
		}
	}
	return nil
}

func (repo *PostgreSQL) CreatePaymentVA(ctx context.Context, orderID string) error {
	// return derrors.NewErrorf(derrors.ErrorCodeUnknown, "Failed to call API Payment Gateway")
	return nil
}

func (repo *PostgreSQL) DecreaseStockProduct(ctx context.Context, tx pgx.Tx, productQuantities map[string]int) error {
	for productID, quantity := range productQuantities {
		query := "UPDATE products SET stock = stock - $1 WHERE id = $2"
		_, err := tx.Exec(ctx, query, quantity, productID)
		if err != nil {
			return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, PostgreErrMsg)
		}
	}
	return nil
}

func (repo *PostgreSQL) DeleteCartItems(ctx context.Context, tx pgx.Tx, cartItemIds []string) error {
	query := `UPDATE cart_items SET deleted_at = now() WHERE id = ANY($1)`
	_, err := tx.Exec(ctx, query, cartItemIds)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, PostgreErrMsg)
	}
	return nil
}
