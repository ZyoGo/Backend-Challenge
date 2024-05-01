package repository

import (
	"github.com/ZyoGo/Backend-Challange/internal/users/core"
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID          string
	Username    string
	Email       string
	Password    string
	Address     string
	PhoneNumber string
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

func FromCore(params core.User) User {
	return User{
		ID:          params.ID,
		Username:    params.Username,
		Email:       params.Email,
		Password:    params.Password,
		Address:     params.Address,
		PhoneNumber: params.PhoneNumber,
	}
}

func (row *User) ToCore() core.User {
	return core.User{
		ID:          row.ID,
		Username:    row.Username,
		Email:       row.Email,
		Password:    row.Password,
		Address:     row.Address,
		PhoneNumber: row.PhoneNumber,
		CreatedAt:   row.CreatedAt.Time,
		UpdatedAt:   row.UpdatedAt.Time,
	}
}
