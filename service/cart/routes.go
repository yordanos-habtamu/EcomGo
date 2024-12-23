package cart

import (
	"fmt"
	"log"
	"net/http"
    "context"
	"github.com/gorilla/mux"
	"github.com/yordanos-habtamu/EcomGo.git/service/auth"
	"github.com/yordanos-habtamu/EcomGo.git/types"
	"github.com/yordanos-habtamu/EcomGo.git/utils"
)

type contextKey string
const myuser contextKey = "user"

type Handler struct {
	store types.OrderStore
	productStore types.ProductStore
	userStore types.UserStore
}

func NewHandler (store types.OrderStore,productStore types.ProductStore,userStore types.UserStore) *Handler{
	return &Handler{store:store,productStore: productStore,userStore: userStore}
}

func (h *Handler) RegisterRoutes (router *mux.Router){
	router.HandleFunc("/cart/checkout",auth.WithJwtAuth(h.handleCheckout,h.userStore)).Methods(http.MethodPost)
}

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request){
	var userID int = 0
	var cart types.CartCheckoutPayload
	if err := utils.ParseJson(r,&cart); err !=nil {
	utils.WriteError(w,http.StatusBadRequest,err)
	return
	}
	if err := utils.Validate.Struct(cart); err != nil {
		log.Printf("Validation error: %v", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}

	productIDs, err:= getCartItemsIDs(cart.Items)
	if err != nil {
		utils.WriteError(w,http.StatusBadRequest,err)
	}
   ps,err := h.productStore.GetProductByIds(productIDs)
   if err != nil {
	utils.WriteError(w,http.StatusInternalServerError,err)
   }
   orderId, totalPrice,err := h.CreateOrder(ps,cart.Items,userID)
   if err != nil {
	utils.WriteError(w,http.StatusInternalServerError,err)
   }
   utils.WriteJson(w,http.StatusOK,map[string] any{
	   "order_id":orderId,
	   "total_price":totalPrice,
   })
}  