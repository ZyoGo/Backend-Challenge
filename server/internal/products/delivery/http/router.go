package http

import "github.com/gorilla/mux"

func RegisterPath(r *mux.Router, h *Handler) {
	if h == nil {
		panic("Products controller cannot be nil")
	}

	r.HandleFunc("/v1/products", h.GetProducts).Methods("GET")
}
