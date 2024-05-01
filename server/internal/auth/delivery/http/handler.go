package http

import (
	"encoding/json"
	"net/http"

	"github.com/ZyoGo/Backend-Challange/internal/auth/core"
	"github.com/ZyoGo/Backend-Challange/internal/auth/delivery/http/request"
	"github.com/ZyoGo/Backend-Challange/internal/auth/delivery/http/response"

	common "github.com/ZyoGo/Backend-Challange/pkg/http"
)

type Handler struct {
	business core.Business
}

func NewHandler(business core.Business) *Handler {
	return &Handler{business}
}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var request request.LoginUserReq

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.NewBadRequestResponse())
		return
	}

	dto := LoginUserDTO(request)
	user, err := h.business.LoginUser(ctx, dto)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp := response.NewLoginResp(user)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
