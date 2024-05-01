package core

import "github.com/ZyoGo/Backend-Challange/pkg/jwt"

type JWT interface {
	GenerateTokenJWT(email, userId string) (jwt.TokenJWT, error)
	ParseAndVerifyJWT(jwtToken string) (jwt.AuthGuardJWT, error)
}
