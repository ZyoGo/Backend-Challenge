package http

import (
	"net/http"

	"github.com/ZyoGo/Backend-Challange/pkg/http/middleware/authguard"
	"github.com/gorilla/mux"
)

func RegisterPath(r *mux.Router, h *Handler, authGuard *authguard.AuthGuard) {
	if h == nil {
		panic("Auth controller cannot be nil")
	}

	r.HandleFunc("/v1/products/cart", h.AddCartItem).Methods("POST").Handler(authGuard.Guard(http.HandlerFunc(h.AddCartItem)))
}