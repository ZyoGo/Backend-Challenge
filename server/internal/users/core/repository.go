package core

import "context"

type Repository interface {
	FindUserByEmail(ctx context.Context, email string) (User, error)
}
