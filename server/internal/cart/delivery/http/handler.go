package http

import (
	"encoding/json"
	"net/http"

	"github.com/ZyoGo/Backend-Challange/internal/cart/core"
	"github.com/ZyoGo/Backend-Challange/internal/cart/delivery/http/request"
	"github.com/ZyoGo/Backend-Challange/internal/cart/delivery/http/response"

	common "github.com/ZyoGo/Backend-Challange/pkg/http"
	jwt "github.com/ZyoGo/Backend-Challange/pkg/jwt"
)

type Handler struct {
	business core.Business
}

func NewHandler(business core.Business) *Handler {
	return &Handler{business}
}

func (h *Handler) AddCartItem(w http.ResponseWriter, r *http.Request) {
	request := new(request.AddCartItemRequest)
	ctx := r.Context()
	user, ok := ctx.Value("userAttr").(jwt.AuthGuardJWT)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(common.NewUnauthorizedResponse("Invalid / expired token"))
	}
	request.UserID = user.UserId

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.NewBadRequestResponse())
		return
	}

	dto := NewAddCartItemDTO(request)
	if err := h.business.AddCartItem(ctx, dto); err != nil {
		resp := common.MapErrorToResponse(err)
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(common.NewSuccessCreatedResponse())
}

func (h *Handler) GetCartByUserID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, ok := ctx.Value("userAttr").(jwt.AuthGuardJWT)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(common.NewUnauthorizedResponse("Invalid / expired token"))
	}
	userID := user.UserId

	carts, err := h.business.GetCartItems(ctx, userID)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		w.WriteHeader(resp.Code)
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp := response.NewGetCartByUserIDResp(carts)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
