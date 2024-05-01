package authguard

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	common "github.com/ZyoGo/Backend-Challange/pkg/http"
	utils "github.com/ZyoGo/Backend-Challange/pkg/jwt"
)

type JWT interface {
	GenerateTokenJWT(email, userId string) (utils.TokenJWT, error)
	ParseAndVerifyJWT(jwtToken string) (utils.AuthGuardJWT, error)
}

type AuthGuard struct {
	jwt JWT
}

func NewBusiness(jwt JWT) *AuthGuard {
	return &AuthGuard{jwt}
}

const (
	// Header
	prefixAuthHeader string = "Bearer "

	// Response Msg
	AuthorizationMsg string = "Authorization header missing/invalid"
	TokenInvalidMsg  string = "Invalid / expired token"

	userAttr string = "userAttr"
)

func (g *AuthGuard) Guard(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if !strings.HasPrefix(authHeader, prefixAuthHeader) {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(common.NewUnauthorizedResponse(AuthorizationMsg))
			return
		}

		attr, err := g.jwt.ParseAndVerifyJWT(strings.TrimPrefix(authHeader, prefixAuthHeader))
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(common.NewUnauthorizedResponse(TokenInvalidMsg))
			return
		}

		// set attr with context
		ctx := r.Context()
		ctx = context.WithValue(ctx, userAttr, attr)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
