package http

import "github.com/gorilla/mux"

func RegisterPath(r *mux.Router, h *Handler) {
	if h == nil {
		panic("Auth controller cannot be nil")
	}

	r.HandleFunc("/v1/auth/login", h.LoginUser).Methods("POST")
}
