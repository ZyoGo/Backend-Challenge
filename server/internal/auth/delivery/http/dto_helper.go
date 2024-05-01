package http

import (
	"github.com/ZyoGo/Backend-Challange/internal/auth/core"
	"github.com/ZyoGo/Backend-Challange/internal/auth/delivery/http/request"
)

func LoginUserDTO(req request.LoginUserReq) core.LoginUserParams {
	return core.LoginUserParams{
		Email:    req.Email,
		Password: req.Password,
	}
}
