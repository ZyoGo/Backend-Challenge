package http

import (
	"encoding/json"
	"net/http"

	"github.com/ZyoGo/Backend-Challange/internal/products/core"
	"github.com/ZyoGo/Backend-Challange/internal/products/delivery/http/request"
	"github.com/ZyoGo/Backend-Challange/internal/products/delivery/http/response"

	common "github.com/ZyoGo/Backend-Challange/pkg/http"
)

type Handler struct {
	business core.Business
}

func NewHandler(business core.Business) *Handler {
	return &Handler{business}
}

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	params := new(request.GetProductsParams)

	ctx := r.Context()
	query := r.URL.Query()
	params.CategoryID = query.Get("category_id")

	dto := NewGetProductsDTO(params)
	products, err := h.business.GetProducts(ctx, dto)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp := response.NewGetProductsResp(products)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
