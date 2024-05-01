package business

import (
	"context"

	"github.com/ZyoGo/Backend-Challange/internal/users/core"
)

type UserBusiness struct {
	repo core.Repository
	id   core.Id
}

func NewBusiness(repo core.Repository, id core.Id) core.Business {
	return &UserBusiness{repo, id}
}

func (s *UserBusiness) GetUserByEmail(ctx context.Context, email string) (core.User, error) {
	return s.repo.FindUserByEmail(ctx, email)
}
