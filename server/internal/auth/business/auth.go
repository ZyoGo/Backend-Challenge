package business

import (
	"context"

	"github.com/ZyoGo/Backend-Challange/internal/auth/core"
	"github.com/ZyoGo/Backend-Challange/pkg/derrors"
)

type AuthBusiness struct {
	userIntraprocess core.UserIntraprocess
	jwt              core.JWT
	hash             core.Hash
}

func NewBusiness(userIntraprocess core.UserIntraprocess, jwt core.JWT, hash core.Hash) core.Business {
	return &AuthBusiness{userIntraprocess, jwt, hash}
}

func (b *AuthBusiness) LoginUser(ctx context.Context, dto core.LoginUserParams) (core.Auth, error) {
	user, err := b.userIntraprocess.GetUserByEmail(ctx, dto.Email)
	if err != nil {
		return core.Auth{}, err
	}

	if !b.hash.CompareHashPassword(dto.Password, user.Password) {
		return core.Auth{}, derrors.NewErrorf(derrors.ErrorCodeNotFound, "Username or Password is invalid")
	}

	token, err := b.jwt.GenerateTokenJWT(user.Email, user.ID)
	if err != nil {
		return core.Auth{}, err
	}

	return core.Auth{
		ID:           user.ID,
		Username:     user.Username,
		Email:        user.Email,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiredAt:    token.ExpiredAt,
	}, nil
}
