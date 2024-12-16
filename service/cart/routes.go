package cart

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yordanos-habtamu/EcomGo.git/types"
	"github.com/yordanos-habtamu/EcomGo.git/utils"
)


type Handler struct {
	store types.OrderStore
}

func NewHandler (store types.OrderStore) *Handler{
	return &Handler{store:store}
}

func (h *Handler) RegisterRoutes (router *mux.Router){
	router.HandleFunc("/cart/checkout",h.handleCheckout).Methods(http.MethodPost)
}

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request){
	var cart types.CartCheckoutPayload
	if err := utils.ParseJson(r,&cart); err !=nil {
	utils.WriteError(w,http.StatusBadRequest,err)
	return
	}
}