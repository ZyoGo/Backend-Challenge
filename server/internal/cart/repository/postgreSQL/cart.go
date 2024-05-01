package postgresql

import (
	"context"

	"github.com/ZyoGo/Backend-Challange/internal/cart/core"
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

func (repo *PostgreSQL) CheckCartByUserID(ctx context.Context, userID string) (exist bool, id string, err error) {
	query := `SELECT id FROM carts WHERE user_id = $1`

	err = repo.db.QueryRow(ctx, query, userID).Scan(&id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return false, id, nil
		}
		return false, id, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, PostgreErrMsg)
	}

	return true, id, nil
}

func (repo *PostgreSQL) CreateCart(ctx context.Context, userID string, newID string) error {
	query := `INSERT INTO carts(id, user_id, created_at, updated_at)
				VALUES($1, $2, now(), now())`

	_, err := repo.db.Exec(ctx, query, newID, userID)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, PostgreErrMsg)
	}

	return nil
}

func (repo *PostgreSQL) InsertOrUpdateCartItem(ctx context.Context, params core.AddCartItemParams) error {
	// Check existing row and get quantity
	var existQuantity int
	err := repo.db.QueryRow(ctx, `SELECT quantity FROM cart_items WHERE cart_id = $1 AND product_id = $2`, params.CartID, params.ProductID).Scan(&existQuantity)
	if err != nil && err != pgx.ErrNoRows {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, PostgreErrMsg)
	}
	// If no rows, insert new data
	if err == pgx.ErrNoRows {
		_, err = repo.db.Exec(ctx, `INSERT INTO cart_items(id, cart_id, product_id, quantity, created_at, updated_at)
		VALUES($1, $2, $3, $4, now(), now())`, params.CartItemID, params.CartID, params.ProductID, params.Quantity)
		if err != nil {
			return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, PostgreErrMsg)
		}
		return nil
	}

	// Update quantity for row
	newQuantity := existQuantity + params.Quantity
	_, err = repo.db.Exec(ctx, `UPDATE cart_items SET quantity = $1, updated_at = now() WHERE cart_id = $2 AND product_id = $3`, newQuantity, params.CartID, params.ProductID)
	return err
}

func (repo *PostgreSQL) FindCartItems(ctx context.Context, cartID string) (core.Cart, error) {
	return core.Cart{}, nil
}
