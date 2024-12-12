package product

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yordanos-habtamu/EcomGo.git/service/middleware"
	"github.com/yordanos-habtamu/EcomGo.git/types"
	"github.com/yordanos-habtamu/EcomGo.git/utils"
)


type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler{
	return &Handler{store:store}
}

// RegisterRoutes registers the routes for product-related operations
func (h *Handler) RegisterRoutes(router *mux.Router) {
	productMiddleware:= middleware.JwtMiddleware("admin","seller")
	adminOnlyMiddleware:= middleware.JwtMiddleware("admin")
	router.HandleFunc("/products",  productMiddleware(h.handleCreateProduct)).Methods("POST")
	router.HandleFunc("/products/{id:[0-9]+}", h.handleGetProductByID).Methods("GET")
	router.HandleFunc("/products", h.handleGetAllProducts).Methods("GET")
	router.HandleFunc("/products/{id:[0-9]+}", adminOnlyMiddleware(h.handleUpdateProduct)).Methods("PUT")
	router.HandleFunc("/products/{id:[0-9]+}", adminOnlyMiddleware( h.handleDeleteProduct)).Methods("DELETE")
	router.HandleFunc("/products/category/{category}", h.handleGetProductsByCategory).Methods("GET")
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter,r *http.Request){
	var payload types.RegisterProductPayload
	if err := utils.ParseJson(r,&payload); err !=nil{
		log.Printf("Failed to parse JSON: %v", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid request format"))
		return
	}
	if err := utils.Validate.Struct(payload); err != nil {
		log.Printf("Validation error: %v", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}
   err := h.store.CreateProduct(types.Product{
		Name:        payload.Name,
		Description: payload.Description,
		Price:       payload.Price,
		Stock:       payload.Stock,
		Catagory:    payload.Catagory,
		ImgUrl:      payload.ImgUrl,
	})
	if err != nil {
		log.Printf("Error creating product : %v",err)
		http.Error(w,"Error creating user",http.StatusInternalServerError)
	    return
		} 

		utils.WriteJson(w,http.StatusCreated,map[string]string{"message":"product created successfully"})


}

func (h *Handler) handleGetProductByID(w http.ResponseWriter,r *http.Request){
	var payload types.RegisterProductPayload
	if err := utils.ParseJson(r,&payload); err !=nil{
		log.Printf("Failed to parse JSON: %v", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid request format"))
		return
	}
	if err := utils.Validate.Struct(payload); err != nil {
		log.Printf("Validation error: %v", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}
	_,err := h.store.GetProductById(payload.ID)


}

func (h *Handler) handleGetAllProducts(w http.ResponseWriter,r *http.Request){
	var payload types.RegisterProductPayload
	if err := utils.ParseJson(r,&payload); err !=nil{
		log.Printf("Failed to parse JSON: %v", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid request format"))
		return
	}
	if err := utils.Validate.Struct(payload); err != nil {
		log.Printf("Validation error: %v", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}


}

func (h *Handler) handleUpdateProduct(w http.ResponseWriter,r *http.Request){
	var payload types.RegisterProductPayload
	if err := utils.ParseJson(r,&payload); err !=nil{
		log.Printf("Failed to parse JSON: %v", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid request format"))
		return
	}
	if err := utils.Validate.Struct(payload); err != nil {
		log.Printf("Validation error: %v", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}


}

func (h *Handler) handleDeleteProduct(w http.ResponseWriter,r *http.Request){
	var payload types.RegisterProductPayload
	if err := utils.ParseJson(r,&payload); err !=nil{
		log.Printf("Failed to parse JSON: %v", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid request format"))
		return
	}
	if err := utils.Validate.Struct(payload); err != nil {
		log.Printf("Validation error: %v", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}


}

func (h *Handler) handleGetProductsByCategory(w http.ResponseWriter,r *http.Request){
	var payload types.RegisterProductPayload
	if err := utils.ParseJson(r,&payload); err !=nil{
		log.Printf("Failed to parse JSON: %v", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid request format"))
		return
	}
	if err := utils.Validate.Struct(payload); err != nil {
		log.Printf("Validation error: %v", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}


}