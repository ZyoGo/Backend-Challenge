package core

import (
	"context"

	"github.com/ZyoGo/Backend-Challange/internal/users/core"
)

type UserIntraprocess interface {
	GetUserByEmail(ctx context.Context, email string) (core.User, error)
}
