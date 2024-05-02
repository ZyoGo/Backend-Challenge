package postgresql

import (
	"context"

	"github.com/ZyoGo/Backend-Challange/internal/carts/core"
	"github.com/ZyoGo/Backend-Challange/pkg/derrors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
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
	query := `SELECT id FROM carts WHERE user_id = $1 AND deleted_at IS NULL`

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
	err := repo.db.QueryRow(ctx, `SELECT quantity FROM cart_items WHERE cart_id = $1 AND product_id = $2 AND deleted_at IS NULL`, params.CartID, params.ProductID).Scan(&existQuantity)
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
	_, err = repo.db.Exec(ctx, `UPDATE cart_items SET quantity = $1, updated_at = now() WHERE cart_id = $2 AND product_id = $3 AND deleted_at IS NULL`, newQuantity, params.CartID, params.ProductID)
	return err
}

func (repo *PostgreSQL) FindCartItems(ctx context.Context, userID string) (core.Cart, error) {
	var (
		carts         core.Cart
		cartID        string
		cartCreatedAt pgtype.Timestamp
		cartUpdatedAt pgtype.Timestamp
	)
	tempCartItems := make(map[string][]core.CartItem)

	query := `SELECT
					c.id as cart_id, c.user_id, c.created_at as cart_created_at, c.updated_at as cart_updated_at,
					ci.id as item_id, ci.product_id, ci.quantity, ci.created_at as item_created_at, ci.updated_at as item_updated_at,
					p.name as product_name, p.price as product_price
			  FROM
					carts c
			  JOIN
					cart_items ci ON c.id = ci.cart_id
			  JOIN
					products p ON ci.product_id = p.id
			  WHERE
					c.user_id = $1 AND ci.deleted_at IS NULL`

	rows, err := repo.db.Query(ctx, query, userID)
	if err != nil {
		return core.Cart{}, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, PostgreErrMsg)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			itemID        string
			productID     string
			quantity      int
			itemCreatedAt pgtype.Timestamp
			itemUpdatedAt pgtype.Timestamp
			productName   string
			productPrice  float64
		)

		err := rows.Scan(&cartID, &userID, &cartCreatedAt, &cartUpdatedAt, &itemID, &productID, &quantity, &itemCreatedAt, &itemUpdatedAt, &productName, &productPrice)
		if err != nil {
			return core.Cart{}, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, PostgreErrMsg)
		}

		if _, ok := tempCartItems[cartID]; !ok {
			tempCartItems[cartID] = []core.CartItem{}
		}

		item := CartItem{
			ID:           itemID,
			ProductID:    productID,
			ProductName:  productName,
			ProductPrice: productPrice,
			Quantity:     quantity,
			CreatedAt:    itemCreatedAt,
			UpdatedAt:    itemUpdatedAt,
		}
		tempCartItems[cartID] = append(tempCartItems[cartID], item.ToCore())
	}

	carts.ID = cartID
	carts.UserID = userID
	carts.CartItem = tempCartItems[cartID]
	carts.CreatedAt = cartCreatedAt.Time
	carts.UpdatedAt = cartUpdatedAt.Time
	return carts, nil
}

func (repo *PostgreSQL) DeleteCartItemByID(ctx context.Context, cartItemID string) error {
	query := `UPDATE cart_items SET deleted_at = now() WHERE id = $1`

	_, err := repo.db.Exec(ctx, query, cartItemID)
	if err != nil {
		return derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, PostgreErrMsg)
	}
	return nil
}
