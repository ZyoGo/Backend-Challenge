package intraprocess

import (
	"context"

	"github.com/ZyoGo/Backend-Challange/internal/users/core"
)

type UserIntraprocess struct {
	business core.Business
}

func NewIntraprocess(business core.Business) UserIntraprocess {
	return UserIntraprocess{business}
}

func (i UserIntraprocess) GetUserByEmail(ctx context.Context, email string) (core.User, error) {
	return i.business.GetUserByEmail(ctx, email)
}
