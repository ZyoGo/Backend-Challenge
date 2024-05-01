package repository

import (
	"context"

	"github.com/ZyoGo/Backend-Challange/internal/users/core"
	"github.com/ZyoGo/Backend-Challange/pkg/derrors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const PostgreErrMsg string = "Postgre error"

type postgreSQL struct {
	db *pgxpool.Pool
}

func NewPostgreSQL(db *pgxpool.Pool) *postgreSQL {
	return &postgreSQL{db}
}

func (repo *postgreSQL) FindUserByEmail(ctx context.Context, email string) (core.User, error) {
	var user User
	query := `SELECT 
					id, username, email, password, 
					address, phone_number, 
					created_at, updated_at 
				FROM users 
				WHERE email = $1`

	err := repo.db.QueryRow(ctx, query, email).Scan(
		&user.ID, &user.Username, &user.Email,
		&user.Password, &user.Address, &user.PhoneNumber,
		&user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return user.ToCore(), derrors.WrapErrorf(err, derrors.ErrorCodeNotFound, "Data not found")
		}

		return user.ToCore(), derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, PostgreErrMsg)
	}

	return user.ToCore(), nil
}
