package http

import (
	"encoding/json"
	"net/http"

	"github.com/ZyoGo/Backend-Challange/internal/orders/core"
	"github.com/ZyoGo/Backend-Challange/internal/orders/delivery/http/request"

	common "github.com/ZyoGo/Backend-Challange/pkg/http"
	jwt "github.com/ZyoGo/Backend-Challange/pkg/jwt"
)

type Handler struct {
	business core.Business
}

func NewHandler(business core.Business) *Handler {
	return &Handler{business}
}

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	request := new(request.CreateOrderRequest)

	user, ok := ctx.Value("userAttr").(jwt.AuthGuardJWT)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(common.NewUnauthorizedResponse("Invalid / expired token"))
		return
	}
	request.UserID = user.UserId

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.NewBadRequestResponse())
		return
	}

	dto := NewCreateOrderDTO(request)
	if err := h.business.CreateOrder(ctx, dto); err != nil {
		resp := common.MapErrorToResponse(err)
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(common.NewSuccessCreatedResponse())
}
