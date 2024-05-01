package core

import "context"

type Business interface {
	GetUserByEmail(ctx context.Context, email string) (User, error)
}
