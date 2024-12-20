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

	r.HandleFunc("/v1/users/{userId}/carts", h.AddCartItem).Methods("POST").Handler(authGuard.Guard(http.HandlerFunc(h.AddCartItem)))
	r.HandleFunc("/v1/users/{userId}/carts", h.GetCartByUserID).Methods("GET").Handler(authGuard.Guard(http.HandlerFunc(h.GetCartByUserID)))
	r.HandleFunc("/v1/users/{userId}/carts/{cartId}", h.DeleteCartItemByID).Methods("DELETE").Handler(authGuard.Guard(http.HandlerFunc(h.DeleteCartItemByID)))
}
