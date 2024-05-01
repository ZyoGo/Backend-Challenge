package response

import (
	"net/http"

	"github.com/ZyoGo/Backend-Challange/internal/auth/core"
)

type Login struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiredAt    int64  `json:"expired_at"`
}

type LoginResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Payload Login  `json:"payload"`
}

func NewLoginResp(result core.Auth) *LoginResp {
	payload := Login{
		ID:           result.ID,
		Username:     result.Username,
		Email:        result.Email,
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
		ExpiredAt:    result.ExpiredAt,
	}

	return &LoginResp{
		Code:    http.StatusOK,
		Message: "SUCCESS",
		Payload: payload,
	}
}
